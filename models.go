package main

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

type Orders struct {
	ID      uint   `json:"order_id"`
	CustID  uint   `json:"order_cust_id"`
	Details string `json:"order_dtl"`
}
