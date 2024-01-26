/* File: employee.go

Description:
 Specification of the Employee struct to define the json data structure.
*/

package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type Employee struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FName   string        `json:"fname" bson:"fname"`
	LName   string        `json:"lname" bson:"lname"`
	Company string        `json:"company" bson:"company"`
	EmpID   string        `json:"emp_id" bson:"emp_id"`
}
