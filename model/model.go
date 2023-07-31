package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func LoadUsersJson() Users {
	userscontent, usercontenterr := ioutil.ReadFile("./users.json")
	if usercontenterr != nil {
		log.Fatal("Error when opening file: ", usercontenterr)
	}
	payload := Users{}
	usercontenterr = json.Unmarshal(userscontent, &payload)
	if usercontenterr != nil {
		log.Fatal("Error during Unmarshal(): ", usercontenterr)
	}
	return payload
}

func LoadTicketJson() UsersTickets {
	ticketscontent, ticketcontenterr := ioutil.ReadFile("./tickets.json")
	if ticketcontenterr != nil {
		log.Fatal("Erro when opening file: ", ticketcontenterr)
	}
	var info UsersTickets
	ticketcontenterr = json.Unmarshal(ticketscontent, &info)
	if ticketcontenterr != nil {
		log.Fatal("Error during Unmarshal(): ", ticketcontenterr)
	}
	return info
}
