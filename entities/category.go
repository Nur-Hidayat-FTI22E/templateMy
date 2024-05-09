package entities

import "time"

// Define your struct here
type NameYourStruct struct{
	// use the same data type as in your database
	Id			uint
	Name		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}