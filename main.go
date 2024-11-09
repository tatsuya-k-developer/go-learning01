package main

import (
	"go-learning01/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/public", public_route)

	// private
	// r.Use(auth.JwtMiddleWare())
	r.GET("/private", private_route)
	// r.Use(auth.AuthWithRole("ADMIN")).GET("/admin", admin_route)

	route.RegisterItemRoute(r)

	r.Run("localhost:8081")
}

func public_route(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "public!!"})
}

func private_route(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "private!!"})
}

func admin_route(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "admin route"})
}
