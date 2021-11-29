package models

import (
	"fmt"
	"time"
)

var Meals = []*Meal{
	{
		Id:          0,
		Name:        "grilled goose",
		Price:       10.3,
		LastOrdered: time.Now(),
	},
	{
		Id:          1,
		Name:        "fried chicken",
		Price:       11.3,
		LastOrdered: time.Now(),
	},
}

var Users = []*User{
	{
		UserID:  0,
		Name:    "gs",
		Balance: 0.0,
	},
	{
		UserID:  1,
		Name:    "jzh",
		Balance: 0.0,
	},
	{
		UserID:  2,
		Name:    "zcm",
		Balance: 0.0,
	},
}

var Orders = []*Order{
	{
		OrderID:         0,
		OrderTime:       time.Now(),
		RequesterName:   Users[0].Name,
		AcceptorName:    Users[1].Name,
		OrderedMealName: Meals[0].Name,
		RequesterId:     0,
		AcceptorId:      1,
		OrderedMealId:   0,
		IsDone:          false,
	},
	{
		OrderID:         1,
		OrderTime:       time.Now(),
		RequesterName:   Users[1].Name,
		AcceptorName:    Users[2].Name,
		OrderedMealName: Meals[1].Name,
		RequesterId:     1,
		AcceptorId:      2,
		OrderedMealId:   1,
		IsDone:          false,
	},
	{
		OrderID:         2,
		OrderTime:       time.Now(),
		RequesterName:   Users[0].Name,
		AcceptorName:    Users[2].Name,
		OrderedMealName: Meals[1].Name,
		RequesterId:     0,
		AcceptorId:      2,
		OrderedMealId:   1,
		IsDone:          false,
	},
}

var console_color = map[string][2]int{
	"order": {31, 40}, //red
	"meal":  {32, 40}, //green
	"user":  {33, 40}, //yellow
}

func PaintStringFunc(obj string) func(fmt.Stringer) string {
	return func(s fmt.Stringer) string {
		return fmt.Sprintf("\033[1;%d;%dm%s\033[0m", console_color[obj][0],
			console_color[obj][1], s)
	}
}