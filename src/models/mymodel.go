package mymodel

import (
	"time"
)

type Model struct {
	FirstBeautifulField string    		`json:"first_beautiful_field" bson:"first_beautiful_field"`
	SecondBeautifulField time.Time 		`json:"second_beautiful_field" bson:"second_beautiful_field"`
}