package model

import (
	"time"
)

type Users []struct {
	Id         int `json:"_id"`
	Name       string
	Created_at time.Time
	Verified   bool
	info       *UsersTickets
}
type UsersTickets []struct {
	Id          string `json:"_id"`
	created_at  time.Time
	Type        string
	Subject     string
	Assignee_id int
	Tags        []string
}
