package transactionsdto

import "time"

type TransactionRequest struct {
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	Price     int    `json:"price"`
}

type CreateTransactionRequest struct {
	StartDate time.Time `json:"startdate"`
	DueDate   time.Time `json:"duedate"`
	UserID    int       `json:"user_id" form:"user_id"`
	Attache   string    `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool      `json:"status" gorm:"type:text" form:"status"`
}

type UpdateTransactionRequest struct {
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	Price     int    `json:"price"`
}
