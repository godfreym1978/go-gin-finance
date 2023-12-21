# go-gin-finance

This sample application uses go gin framework to build a REST API application that provides access to a backend database

It runs on the machine in port 8080.

It needs a MySQL database with the orders table created. The script for the table is - 

CREATE TABLE orders{
    order_id int,
    order_cust_id int,
    order_dtl varchar(100)
}

Populate the table with the following values

insert into orders values(1001,1,'Ship');
insert into orders values(1002,2,'Pen');
insert into orders values(1003,3,'Table');

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

POST - http://localhost:8080/api/v1/orders
It will create a record in the table with the following JSON data
{
    "order_id": 1007,
    "order_cust_id": 3,
    "order_dtl": "pen"
}