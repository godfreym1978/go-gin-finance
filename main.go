package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/orders", GetOrders)

		v1.GET("/orders/:id", GetOrder)
		v1.POST("/orders", CreateOrder)
		/*
			v1.PUT("/users/:id", UpdateUser)
			v1.DELETE("/users/:id", DeleteUser)

			v1.GET("/posts", GetPosts)
			v1.GET("/posts/:id", GetPost)
			v1.POST("/posts", CreatePost)
			v1.PUT("/posts/:id", UpdatePost)
			v1.DELETE("/posts/:id", DeletePost)
		*/
	}

	router.Run(":8080")
}
