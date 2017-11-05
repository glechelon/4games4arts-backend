package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/glechelon/4games4arts/data"
	"github.com/glechelon/4games4arts/handlers"
)

func main() {

	myPost := &data.Post{}
	myPost.Id = "5"
	fmt.Println(myPost.Id)

	router := gin.Default()
	router.POST("/post", handlers.CreatePost)
	router.GET("/post", handlers.FetchAllPosts)
	router.GET("/post/:id", handlers.FetchPost)
	router.PUT("/post/:id", handlers.UpdatePost)
	router.DELETE("/post/:id", handlers.DeletePost)

	router.Run()
}
