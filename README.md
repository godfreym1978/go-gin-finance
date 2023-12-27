# go-gin-finance

This sample application uses go gin framework to build a REST API application that provides access to a backend database

It runs on the machine in port 8080.

It needs a MySQL database with the orders table. The table wull be created in the database when the application is run the first time around when the main.go executes the dbutils package's "init-table.go" which contains in turn uses the model.go package to execute the creation ot tables in the DB.

Populate the table with POST requests

POST - http://localhost:8080/api/v1/orders
It will create a record in the table with the following JSON data
{
    "order_id": 1007,
    "order_cust_id": 3,
    "order_dtl": "pen"
}

To invoke the services use postman or similar services -

GET - http://localhost:8080/api/v1/orders
It will  return all the rows from the  DB similar to  - 
[
    {
        "order_id": 1001,
        "order_cust_id": 1,
        "order_dtl": "dinner table"
    },
    {
        "order_id": 1002,
        "order_cust_id": 1,
        "order_dtl": "office table"
    }
]

GET - http://localhost:8080/api/v1/orders/1005
It will return a single record from the DB

{
    "order_id": 1005,
    "order_cust_id": 3,
    "order_dtl": "pen"
}

This project also involves the use of MongoDB connector to serach and insert records in the MongoDB DB Collection