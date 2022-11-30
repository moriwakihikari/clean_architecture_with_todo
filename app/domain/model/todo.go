package model

import "time"

type Todo struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Note     string    `json:"note"`
	DayTime  time.Time `json:"day_time"`
	UserID   int       `json:"user_id"`
}

type Todos []Todo