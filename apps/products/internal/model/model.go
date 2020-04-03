package model

import "time"

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64
	Content string
	Author  string
}

type Product struct {
	Detail  *Detail   `json:"detail"`
	Rating  *Rating   `json:"rating"`
	Reviews []*Review `json:"reviews"`
}

type Detail struct {
	ID          int64    `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	CreatedTime time.Time `json:"created_time"`
}

type Rating struct {
	ID          int64    `json:"id"`
	ProductID   int64    `json:"product_id"`
	Score       uint32    `json:"score"`
	UpdatedTime time.Time `json:"updated_time"`
}

type Review struct {
	ID          int64    `json:"id"`
	ProductID   int64    `json:"product_id"`
	Message     string    `json:"message"`
	CreatedTime time.Time `json:"created_time"`
}
