package main

import (
	"go-gin-finance/dbutils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// if there is an error opening the connection, handle it
	db := createDBConn()

	dbutils.Initialize(db)

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
