package dbutils

import (
	"gopkg.in/mgo.v2/bson"
)

type Orders struct {
	ID      uint   `json:"order_id"`
	CustID  uint   `json:"order_cust_id"`
	Details string `json:"order_dtl"`
}

type Employee struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FName   string        `json:"fname" bson:"fname"`
	LName   string        `json:"lname" bson:"lname"`
	Company string        `json:"company" bson:"company"`
	EmpID   string        `json:"emp_id" bson:"emp_id"`
}

const orders = `
CREATE TABLE IF NOT EXISTS orders (
order_id INTEGER ,
order_cust_id INTEGER ,
order_dtl VARCHAR(100) NULL
)
`
