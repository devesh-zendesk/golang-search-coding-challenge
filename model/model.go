package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type User struct {
	Id        int       `json:"_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Verified  bool      `json:"verified"`
	Tickets   []Ticket
}
type Ticket struct {
	Id         string    `json:"_id"`
	CreatedAt  time.Time `json:"created_at"`
	Type       string    `json:"type"`
	Subject    string    `json:"subject"`
	AssigneeId int       `json:"assignee_id"`
	Tags       []string  `json:"tags"`
	Assignee   User
}

type Index struct {
	reversedMapUsersById   map[int]*User
	reversedMapUsersByName map[string]*User
}

type UserFinder interface {
	GetUserById(id int) (User, error)
	GetUserByName(name string) (User, error)
}

func (index *Index) GetUserById(id int) (User, error) {
	user, exists := index.reversedMapUsersById[id]
	if !exists {
		return User{}, fmt.Errorf("User ID %d does not Exist", id)
	}
	return *user, nil
}

func (index *Index) GetUserByName(name string) (User, error) {
	user, exists := index.reversedMapUsersByName[name]
	if !exists {
		return User{}, fmt.Errorf("User with name %s does not Exist", name)
	}
	return *user, nil
}

func Init() UserFinder {

	ticketsContent, err := ioutil.ReadFile("./tickets.json")
	if err != nil {
		log.Fatal("Erro when opening file: ", err)
	}
	var info []Ticket
	if err = json.Unmarshal(ticketsContent, &info); err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	usersContent, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload []User
	if err = json.Unmarshal([]byte(usersContent), &payload); err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	index := Index{
		reversedMapUsersById:   make(map[int]*User),
		reversedMapUsersByName: make(map[string]*User),
	}
	for _, record := range payload {
		id := record.Id
		name := strings.ToLower(record.Name)
		index.reversedMapUsersById[id] = &record
		index.reversedMapUsersByName[name] = &record
	}

	return &index
}
