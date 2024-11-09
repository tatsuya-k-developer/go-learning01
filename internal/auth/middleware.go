package auth

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getKey(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte("unko"), nil
}

func AuthWithRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := ctx.Get("claims")
		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{"msg": "invalid token"})
		} else {
			if claims.(jwt.MapClaims)["role"] == role {
				ctx.Next()
			} else {
				ctx.AbortWithStatusJSON(403, gin.H{"msg": "forbidden"})
			}
		}
	}
}

func JwtMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// sample token string taken from the New example
		tokenString := ctx.GetHeader("X-API-KEY")
		log.Println("key: \"" + tokenString + "\"")

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, getKey)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"msg": "invalid token"})
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				ctx.Set("claims", claims)
			} else {
				log.Println(err)
				ctx.AbortWithStatusJSON(401, gin.H{"msg": "error"})

			}

			ctx.Next()
		}
	}
}
