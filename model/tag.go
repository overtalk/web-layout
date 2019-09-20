package model

type Tag struct {
	Detail string `json:"detail"`
	ID     int    `json:"id;primary_key"`
}
