package model

import "time"

type Detail struct {
	ID    int64 `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	CreatedTime time.Time `json:"created_time"`
}

type Article struct {
	ID int64
	Content string
	Author string
}