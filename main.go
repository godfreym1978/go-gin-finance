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
			v1.PUT("/orders/:id", UpdateUser)
			v1.DELETE("/orders/:id", DeleteUser)

		*/
	}

	router.Run(":8080")
}
