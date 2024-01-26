package dbutils

const orders = `
CREATE TABLE IF NOT EXISTS orders (
order_id INTEGER ,
order_cust_id INTEGER ,
order_dtl VARCHAR(100) NULL
)
`
