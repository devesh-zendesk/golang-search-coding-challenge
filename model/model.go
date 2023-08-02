package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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
	users                    []User
	tickets                  []Ticket
	reversedMapUsersById     map[int]User
	reversedMapUsersByName   map[string][]User
	reversedMapUsersByVerify map[string][]User
}

type UserFinder interface {
	GetUserById(id int) (User, error)
	GetUserByName(name string) ([]User, error)
	GetUserByVerifiedFlag(flag string) ([]User, error)
}

func Init() UserFinder {

	ticketsContent, err := ioutil.ReadFile("./tickets.json")
	if err != nil {
		log.Fatal("Erro when opening file: ", err)
	}
	var tickets []Ticket
	if err = json.Unmarshal(ticketsContent, &tickets); err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	usersContent, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var users []User
	if err = json.Unmarshal([]byte(usersContent), &users); err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	index := Index{
		reversedMapUsersById:     make(map[int]User),
		reversedMapUsersByName:   make(map[string][]User),
		reversedMapUsersByVerify: make(map[string][]User),
		users:                    users,
		tickets:                  tickets,
	}

	index.LoadIndex()

	return &index
}

func (index *Index) GetUserById(id int) (User, error) {
	user, exists := index.reversedMapUsersById[id]
	if !exists {
		return User{}, fmt.Errorf("User ID %d does not Exist", id)
	}
	return user, nil
}

func (index *Index) GetUserByName(name string) ([]User, error) {
	user, exists := index.reversedMapUsersByName[name]
	if !exists {
		return []User{}, fmt.Errorf("User with name %s does not Exist", name)
	}
	return user, nil
}

func (index *Index) GetUserByVerifiedFlag(flag string) ([]User, error) {
	user, exists := index.reversedMapUsersByVerify[flag]
	if !exists {
		return []User{}, fmt.Errorf("User with the flag %s does not Exist", flag)
	}
	return user, nil
}

func (index *Index) LoadIndex() {
	for _, record := range index.users {
		id := record.Id
		name := strings.ToLower(record.Name)
		name_splice := strings.Split(name, " ")
		verify_flag := strconv.FormatBool(record.Verified)
		index.reversedMapUsersById[id] = record
		index.reversedMapUsersByName[name] = append(index.reversedMapUsersByName[name], record)
		for _, item := range name_splice {
			index.reversedMapUsersByName[item] = append(index.reversedMapUsersByName[item], record)
		}
		index.reversedMapUsersByVerify[verify_flag] = append(index.reversedMapUsersByVerify[verify_flag], record)
	}

	for i := 0; i < len(index.tickets); i++ {
		id := index.tickets[i].AssigneeId
		assignee_by_id := index.reversedMapUsersById[id]
		index.tickets[i].Assignee = assignee_by_id
		assignee_by_id.Tickets = append(assignee_by_id.Tickets, index.tickets[i])
		index.reversedMapUsersById[id] = assignee_by_id

		assignee_by_name := index.reversedMapUsersByName[strings.ToLower(assignee_by_id.Name)]
		for j := 0; j < len(assignee_by_name); j++ {
			assignee_by_name[j].Tickets = append(assignee_by_name[j].Tickets, index.tickets[i])
		}
		index.reversedMapUsersByName[strings.ToLower(assignee_by_id.Name)] = assignee_by_name

		name_splice := strings.Split(assignee_by_id.Name, " ")
		for _, name := range name_splice {
			assignee_by_name := index.reversedMapUsersByName[strings.ToLower(name)]
			for j := 0; j < len(assignee_by_name); j++ {
				assignee_by_name[j].Tickets = append(assignee_by_name[j].Tickets, index.tickets[i])
			}
			index.reversedMapUsersByName[strings.ToLower(name)] = assignee_by_name
		}
	}
}
