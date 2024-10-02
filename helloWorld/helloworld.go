package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context)  {
		c.String(200," Hello World")
	})

	api := r.Group("/api")

	api.GET("/teste", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "teste",
		})
	})

	r.Run(":8080")
}