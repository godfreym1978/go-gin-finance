/*
	File: controllers.go

Description:

	This is used to connect to MongoDB and MySQL to query, fetch or insert records.
*/
package main

import (
	"database/sql"
	"fmt"
	"go-gin-finance/config"
	"go-gin-finance/entity"
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

/*

Receives a JSON document in format specified to insert into the MongoDB
{
		{"fname", "John"},
		{"lname", "Doe"},
		{"company", "Great Corporation"},
		{"emp_id", "123456"},
	}

*/

func PutEmployee(c *gin.Context) {
	// Implement logic to get employee from MongoDB

	config, err := config.BuildConfig()
	dsn := fmt.Sprintf("mongodb://%s:%d/", config.Mongo.Host, config.Mongo.Port)
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(c, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(c, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(config.Mongo.Database).Collection(config.Mongo.Collection)

	/* to insert a document into the MongoDB
	var document interface{}
	document = bson.D{
		{"fname", "John"},
		{"lname", "Doe"},
		{"company", "Great Corporation"},
		{"emp_id", "123456"},
	}

	//_, err1 := collection.InsertOne(c, document)

	*/

	//to insert a JSON request into the DB
	var payload entity.Employee
	c.BindJSON(&payload)

	_, err1 := collection.InsertOne(c, payload)

	if err1 != nil {
		fmt.Print(fmt.Errorf("failed with %w", err1).Error())
	}

}

/*
To get all the records from the collection in the MongoDB
*/
func GetEmployees(c *gin.Context) {
	// Implement logic to get employee from MongoDB
	/*
		Receive a JSON request message and pass it onto the search query to get the result.
		Used the following StackOverflow to get help with turning request to JSON query-
		https://stackoverflow.com/questions/39785289/how-to-marshal-json-string-to-bson-document-for-writing-to-mongodb
	*/

	config, err := config.BuildConfig()
	dsn := fmt.Sprintf("mongodb://%s:%d/", config.Mongo.Host, config.Mongo.Port)
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(c, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(config.Mongo.Database).Collection(config.Mongo.Collection)

	var payload []entity.Employee

	cursor, err1 := collection.Find(c, bson.M{})

	if err1 != nil {
		fmt.Print(fmt.Errorf("failed with %w", err1).Error())
	}

	cursor.All(c, &payload)

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, payload)

}

/*
To get records from the collection in the MongoDB using the search criteria as part of JSON post request. The message format can be something like this -
{"fname": "John"}
*/

func GetEmployee(c *gin.Context) {
	// Implement logic to get employee from MongoDB

	config, err := config.BuildConfig()
	dsn := fmt.Sprintf("mongodb://%s:%d/", config.Mongo.Host, config.Mongo.Port)
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(c, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(config.Mongo.Database).Collection(config.Mongo.Collection)

	jsonData, err := io.ReadAll(c.Request.Body)

	var bdoc interface{}
	bson.UnmarshalJSON(jsonData, &bdoc)

	var payload []entity.Employee

	cursor, err1 := collection.Find(c, bdoc)

	if err1 != nil {
		fmt.Print(fmt.Errorf("failed with %w", err1).Error())
	}

	cursor.All(c, &payload)

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, payload)

}

/*
Utility function to initiate the DB connection with MYSQL DB. The format is -

("mysql", "user:password@tcp(host-ip:mysql-port)/db-name")
*/
func createDBConn() *sql.DB {

	config, err := config.BuildConfig()
	fmt.Print(config)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Mysql.User, config.Mysql.Password,
		config.Mysql.Host, config.Mysql.Port, config.Mysql.Database)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	return db
}

/*
Function to get all the records from the table
*/
func GetOrders(c *gin.Context) {
	// if there is an error opening the connection, handle it
	db := createDBConn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	rows, _ := db.Query("select * from orders")

	var orderData []entity.Orders
	var order entity.Orders
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

/*
Function to get a single record from the table-
The GET request will be in the format - http://localhost:8080/api/v1/orders/1007 with 1007 as the ID to be substituted in search in the DB
*/
func GetOrder(c *gin.Context) {
	// Implement logic to fetch a single user

	// if there is an error opening the connection, handle it
	db := createDBConn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	urlPathElements := strings.Split(c.Request.URL.Path, "/")
	var order entity.Orders

	rows := db.QueryRow("select * from orders where order_id = ?", urlPathElements[4])

	err := rows.Scan(&order.ID, &order.CustID, &order.Details)

	if err != nil {
		fmt.Errorf("failed with %w", err).Error()
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, order)

}

/*
Create a record in the DB as a POST. The input record will be in the format -

{
  "fname": "Jane",
  "lname": "Doe",
  "company": "Google",
  "emp_id": "123456"
}
*/

func CreateOrder(c *gin.Context) {
	// if there is an error opening the connection, handle it
	db := createDBConn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var payload entity.Orders

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
