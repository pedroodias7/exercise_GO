package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts []Post

func main() {
	r := gin.Default()

	posts = append(posts, Post{ID: "1", Title: "2", Content: "wwwww"})
	posts = append(posts, Post{ID: "2", Title: "2", Content: "wwwww"})

	r.GET("/tasks", GetTasks)
	r.POST("/tasks", addTask)
	r.PUT("/tasks/:id", updateTaks)
	r.DELETE("tasks/:id", deleteTasks)

	r.Run(":8080")
}

func addTask(c *gin.Context) {

	var newPost Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	newPost.ID = "2"
	newPost.Title = "Teste"
	newPost.Content = "teste"
	fmt.Print("Prind id:", newPost)
	posts = append(posts, newPost)
	// posts = append(posts, newPost)

	c.JSON(http.StatusCreated, newPost)

}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}
func updateTaks(c *gin.Context) {
	id := c.Param("id")

	var updatePost Post
	postFound := false

	// Loop through the posts to find the one with the matching ID
	for i, post := range posts {
		if post.ID == id {
			// Bind the incoming JSON to updatePost
			if err := c.ShouldBindJSON(&updatePost); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error()})
				return
			}

			// Update the post's Title and Content
			posts[i].Title = updatePost.Title
			posts[i].Content = updatePost.Content

			postFound = true
			break
		}
	}

	// If the post was not found, return a 404
	if !postFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}

	// Return the updated post
	c.JSON(http.StatusOK, gin.H{"Message": "Post updated successfully", "Post": updatePost})
}

func deleteTasks(c *gin.Context) {
	id := c.Param("id")

	var deletedPost Post
	postFound := false

	for i, post := range posts {
		if post.ID == id {
			deletedPost = posts[i]
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"Message": "Post delete with suceess"})
			return
		}
		postFound = true
		break
	}

	if !postFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Post deleted", "Post": deletedPost})

}
