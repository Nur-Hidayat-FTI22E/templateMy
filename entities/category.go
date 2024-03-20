package entities

import "time"

type Categories struct{
	Id			uint
	Name		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}