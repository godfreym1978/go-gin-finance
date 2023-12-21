package dbutils

type Orders struct {
	ID      uint   `json:"order_id"`
	CustID  uint   `json:"order_cust_id"`
	Details string `json:"order_dtl"`
}

const orders = `
CREATE TABLE IF NOT EXISTS orders (
order_id INTEGER ,
order_cust_id INTEGER ,
order_dtl VARCHAR(100) NULL
)
`
