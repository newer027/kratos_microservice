package model

import "time"

type Review struct {
	ID          int64    	`json:"id"`
	ProductID   int64    	`json:"product_id"`
	Message     string    	`json:"message"`
	CreatedTime time.Time 	`json:"created_time"`
}