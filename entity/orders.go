/* File: orders.go

Description:
 Specification of the Orders struct to define the json data structure.
*/

package entity

type Orders struct {
	ID      uint   `json:"order_id"`
	CustID  uint   `json:"order_cust_id"`
	Details string `json:"order_dtl"`
}
