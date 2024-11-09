package route

import (
	"go-learning01/internal/repositories"
	"go-learning01/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoute(r *gin.Engine) {
	r.Use(func(ctx *gin.Context) {
		ctx.Set("itemRepository", repositories.NewItemRepository())
	})

	r.Use(func(ctx *gin.Context) {
		if repository, exists := ctx.Get("itemRepository"); exists {
			itemService := service.NewItemService(repository.(repositories.IItemRepository))
			ctx.Set("itemService", itemService)
		}
	})

	router := r.Group("/item")

	router.GET("", getItems)
	router.GET("/:id", getItem)
	// router.POST("/items", CreateItem)
	// router.PUT("/items/:id", UpdateItem)
	// router.DELETE("/items/:id", DeleteItem)
}

func getService(ctx *gin.Context) (service.IItemService, bool) {
	itemService, exists := ctx.Get("itemService")

	if exists {
		return itemService.(service.IItemService), true
	} else {
		return nil, false
	}
}

func getItems(ctx *gin.Context) {
	itemService, exists := getService(ctx)

	if !exists {
		ctx.AbortWithStatusJSON(500, gin.H{"msg": "internal server error"})
	} else {
		items := itemService.GetAll()
		ctx.JSON(200, items)
	}
}

func getItem(ctx *gin.Context) {
	itemService, exists := getService(ctx)

	if !exists {
		ctx.AbortWithStatusJSON(500, gin.H{"msg": "internal server error"})
	} else {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.AbortWithStatusJSON(403, gin.H{"msg": "bad request"})
		}

		item, err := itemService.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(404, gin.H{"msg": "not found"})
		} else {
			ctx.JSON(200, item)
		}
	}
}
