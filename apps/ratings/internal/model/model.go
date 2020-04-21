package model

import "time"

type Rating struct {
	ID          int64    `json:"id"`
	ProductID   int64    `json:"product_id"`
	Score       uint32    `json:"score"`
	UpdatedTime time.Time `json:"updated_time"`
}
