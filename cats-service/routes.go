package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.Use(cors.Default())

	r.GET("/api/meow", GenerateCat)
	r.GET("/api/cats", FindCats)
	r.POST("/api/cats", CreateCat)
	r.GET("/api/cats/:id", FindCat)
	r.PATCH("/cats/:id", UpdateCat)
	r.DELETE("/api/cats/:id", DeleteCat)
	return r
}
