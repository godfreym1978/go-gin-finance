package main

import (
	"database/sql"
	"fmt"
	"go-gin-finance/dbutils"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PutEmployee(c *gin.Context) {
	// Implement logic to get employee from MongoDB

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(c, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testdb").Collection("employee")

	/* to insert a document into the MongoDB
	var document interface{}
	document = bson.D{
		{"fname", "Godfrey"},
		{"lname", "Menezes"},
		{"company", "IBM"},
		{"emp_id", "959454"},
	}

	//_, err1 := collection.InsertOne(c, document)

	*/

	//to insert a JSON request into the DB
	var payload dbutils.Employee
	c.BindJSON(&payload)

	_, err1 := collection.InsertOne(c, payload)

	if err1 != nil {
		fmt.Print(fmt.Errorf("failed with %w", err1).Error())
	}

}

func GetEmployees(c *gin.Context) {
	// Implement logic to get employee from MongoDB
	/*
		Receive a JSON request message and pass it onto the search query to get the result.
		Used the following StackOverflow to get help with turning request to JSON query-
		https://stackoverflow.com/questions/39785289/how-to-marshal-json-string-to-bson-document-for-writing-to-mongodb
	*/

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(c, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testdb").Collection("employee")

	//to insert a JSON request into the DB
	var payload []dbutils.Employee

	cursor, err1 := collection.Find(c, bson.M{})

	if err1 != nil {
		fmt.Print(fmt.Errorf("failed with %w", err1).Error())
	}

	cursor.All(c, &payload)

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, payload)

}

func GetEmployee(c *gin.Context) {
	// Implement logic to get employee from MongoDB

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(c, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testdb").Collection("employee")

	jsonData, err := io.ReadAll(c.Request.Body)

	var bdoc interface{}
	bson.UnmarshalJSON(jsonData, &bdoc)

	var payload []dbutils.Employee

	cursor, err1 := collection.Find(c, bdoc)

	if err1 != nil {
		fmt.Print(fmt.Errorf("failed with %w", err1).Error())
	}

	cursor.All(c, &payload)

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, payload)

}

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

	var orderData []dbutils.Orders
	var order dbutils.Orders
	for rows.Next() {

		err := rows.Scan(&order.ID, &order.CustID, &order.Details)

		if err != nil {
			fmt.Errorf("failed with %w", err).Error()
		}
		orderData = append(orderData, order)

	}
	c.Header("Content-Type", "application/json")
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
	var order dbutils.Orders

	rows := db.QueryRow("select * from orders where order_id = ?", urlPathElements[4])

	err := rows.Scan(&order.ID, &order.CustID, &order.Details)

	if err != nil {
		fmt.Errorf("failed with %w", err).Error()
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, order)
	//c.JSON(201, order)

}

/*
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

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, jsonData)

}
*/

func CreateOrder(c *gin.Context) {
	// if there is an error opening the connection, handle it
	db := createDBConn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var payload dbutils.Orders

	if err := c.BindJSON(&payload); err == nil {

		insertDynStmt := `insert into orders(order_id, order_cust_id, order_dtl) values (?, ?, ?)`

		_, err1 := db.Exec(insertDynStmt, payload.ID, payload.CustID, payload.Details)

		if err1 != nil {
			fmt.Print(fmt.Errorf("error occured during db insert input data %w", err1).Error())
		}

	}

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, payload)

}

/*
func UpdateUser(c *gin.Context) {
	// Implement logic to update an existing user
}

func DeleteUser(c *gin.Context) {
	// Implement logic to delete a user
}

*/
