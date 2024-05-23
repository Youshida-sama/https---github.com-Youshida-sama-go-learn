package models

import "time"

type UserAction struct {
	ID      int
	Name    string
	Surname string
	Time    time.Time
}
