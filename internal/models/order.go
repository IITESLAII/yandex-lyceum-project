package models

type Position struct {
	ID    int    `json:"id" db:"id"`
	Price int    `json:"price" db:"price"`
	Name  string `json:"name" db:"name"`
}
