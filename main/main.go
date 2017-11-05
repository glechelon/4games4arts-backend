package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glechelon/4games4arts-backend/data"
	"github.com/glechelon/4games4arts-backend/handlers"
)

func main() {

	myPost := &data.Post{}
	myPost.Id = "5"
	fmt.Println(myPost.Id)

	router := gin.Default()
	router.Use(cors.Default())

	postGroup := router.Group("/post")

	{
		postGroup.POST("/", handlers.CreatePost)
		postGroup.GET("/", handlers.FetchAllPosts)
		postGroup.GET("/:id", handlers.FetchPost)
		postGroup.PUT("/:id", handlers.UpdatePost)
		postGroup.DELETE("/:id", handlers.DeletePost)
	}

	router.Run()
}
