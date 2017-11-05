package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/glechelon/4games4arts-backend/data"
)

func CreatePost(c *gin.Context) {

}

func FetchAllPosts(c *gin.Context) {
	posts := []*data.Post{&data.Post{
		Id:     "1",
		Date:   "03/11/2017",
		Title:  "post1",
		Author: "jean pierre du chateau bleu",
		Image:  "", Teaser: "this is a teaser",
		Content: "this is a content"}}

	c.JSON(200, posts)
}

func FetchPost(c *gin.Context) {
	post := &data.Post{
		Id:     "1",
		Date:   "03/11/2017",
		Title:  "post1",
		Author: "jean pierre du chateau bleu",
		Image:  "", Teaser: "this is a teaser",
		Content: "this is a content"}

	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {

}

func DeletePost(c *gin.Context) {

}
