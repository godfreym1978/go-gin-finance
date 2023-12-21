package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func createDBConn() *sql.DB {
	db, err := sql.Open("mysql", "root:passw0rd@tcp(127.0.0.1:3306)/finance")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetOrders(c *gin.Context) {
	// if there is an error opening the connection, handle it
	db := createDBConn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	rows, _ := db.Query("select * from orders")

	var orderData []Orders
	var order Orders
	for rows.Next() {

		err := rows.Scan(&order.ID, &order.CustID, &order.Details)

		if err != nil {
			fmt.Errorf("failed with %w", err).Error()
		}
		orderData = append(orderData, order)

	}
	c.IndentedJSON(http.StatusOK, orderData)

}

func GetOrder(c *gin.Context) {
	// Implement logic to fetch a single user

	// if there is an error opening the connection, handle it
	db := createDBConn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	urlPathElements := strings.Split(c.Request.URL.Path, "/")
	//fmt.Print(urlPathElements[4])
	var order Orders

	rows := db.QueryRow("select * from orders where order_id = ?", urlPathElements[4])

	err := rows.Scan(&order.ID, &order.CustID, &order.Details)

	if err != nil {
		fmt.Errorf("failed with %w", err).Error()
	}

	c.IndentedJSON(http.StatusOK, order)

}

func CreateOrder(c *gin.Context) {
	// if there is an error opening the connection, handle it
	db := createDBConn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Errorf("error occured during readin input data %w", err)
	}

	var payload map[string]interface{}

	if err := json.Unmarshal(jsonData, &payload); err != nil {
		fmt.Errorf("error occured during unmarshal data %w", err)
	}

	insertDynStmt := `insert into orders(order_id, order_cust_id, order_dtl) values (?, ?, ?)`

	_, err1 := db.Exec(insertDynStmt, payload["order_id"], payload["order_cust_id"], payload["order_dtl"])

	if err1 != nil {
		fmt.Print(fmt.Errorf("error occured during db insert input data %w", err1).Error())
	}

}

/*
func UpdateUser(c *gin.Context) {
	// Implement logic to update an existing user
}

func DeleteUser(c *gin.Context) {
	// Implement logic to delete a user
}

*/
