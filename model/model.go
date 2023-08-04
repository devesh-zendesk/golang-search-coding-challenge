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
	users                     []User
	tickets                   []Ticket
	reversedMapUsersById      map[int]User
	reversedMapUsersByName    map[string][]User
	reversedMapUsersByVerify  map[string][]User
	reversedMapUsersCreated   map[string][]User
	reversedMapTicketsType    map[string][]Ticket
	reversedMapTicketsTags    map[string][]Ticket
	reversedMapTicketsCreated map[string][]Ticket
}

type UserFinder interface {
	GetUserById(id int) (User, error)
	GetUserByName(name string) ([]User, error)
	GetUserByVerifiedFlag(flag string) ([]User, error)
	GetUserByCreatedDate(date string) ([]User, error)
	GetTicketsByType(ticket_type string) ([]Ticket, error)
	GetTicketsByTag(tag string) ([]Ticket, error)
	GetTicketsByCreatedDate(date string) ([]Ticket, error)
	GetTicketsWithoutAssignee() ([]Ticket, error)
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
		users:                     users,
		tickets:                   tickets,
		reversedMapUsersById:      make(map[int]User),
		reversedMapUsersByName:    make(map[string][]User),
		reversedMapUsersByVerify:  make(map[string][]User),
		reversedMapUsersCreated:   make(map[string][]User),
		reversedMapTicketsType:    make(map[string][]Ticket),
		reversedMapTicketsTags:    make(map[string][]Ticket),
		reversedMapTicketsCreated: make(map[string][]Ticket),
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

func (index *Index) GetUserByCreatedDate(date string) ([]User, error) {
	user, exists := index.reversedMapUsersCreated[date]
	if !exists {
		return []User{}, fmt.Errorf("User with Created Date %s does not Exist", date)
	}
	return user, nil
}

func (index *Index) GetTicketsByType(ticket_type string) ([]Ticket, error) {
	ticket, exists := index.reversedMapTicketsType[ticket_type]
	if !exists {
		return []Ticket{}, fmt.Errorf("Ticket with Type %s does not Exist", ticket_type)
	}
	return ticket, nil
}

func (index *Index) GetTicketsByTag(tag string) ([]Ticket, error) {
	ticket, exists := index.reversedMapTicketsTags[tag]
	if !exists {
		return []Ticket{}, fmt.Errorf("tickets with tag %s does not exist", tag)
	}
	return ticket, nil
}

func (index *Index) GetTicketsByCreatedDate(date string) ([]Ticket, error) {
	ticket, exists := index.reversedMapTicketsCreated[date]
	if !exists {
		return []Ticket{}, fmt.Errorf("tickets with Created Date %s does not exist", date)
	}
	return ticket, nil
}

func (index *Index) GetTicketsWithoutAssignee() ([]Ticket, error) {
	var tickets []Ticket
	for _, ticket := range index.tickets {
		if ticket.AssigneeId == 0 {
			tickets = append(tickets, ticket)
		}
	}
	return tickets, nil
}

func (index *Index) LoadIndex() {
	for _, record := range index.users {
		id := record.Id
		name := strings.ToLower(record.Name)
		name_splice := strings.Split(name, " ")
		verify_flag := strconv.FormatBool(record.Verified)
		formated_date := record.CreatedAt.Format("2006-01-02")
		index.reversedMapUsersById[id] = record
		index.reversedMapUsersByName[name] = append(index.reversedMapUsersByName[name], record)
		for _, item := range name_splice {
			index.reversedMapUsersByName[item] = append(index.reversedMapUsersByName[item], record)
		}
		index.reversedMapUsersByVerify[verify_flag] = append(index.reversedMapUsersByVerify[verify_flag], record)
		index.reversedMapUsersCreated[formated_date] = append(index.reversedMapUsersCreated[formated_date], record)
	}

	for i := 0; i < len(index.tickets); i++ {
		id := index.tickets[i].AssigneeId
		assignee_by_id := index.reversedMapUsersById[id]
		index.tickets[i].Assignee = assignee_by_id
		assignee_by_id.Tickets = append(assignee_by_id.Tickets, index.tickets[i])
		index.reversedMapUsersById[id] = assignee_by_id

		ticket_type := index.tickets[i].Type
		index.reversedMapTicketsType[ticket_type] = append(index.reversedMapTicketsType[ticket_type], index.tickets[i])

		for _, tag := range index.tickets[i].Tags {
			index.reversedMapTicketsTags[strings.ToLower(tag)] = append(index.reversedMapTicketsTags[strings.ToLower(tag)], index.tickets[i])
		}

		ticket_date := index.tickets[i].CreatedAt.Format("2006-01-02")
		index.reversedMapTicketsCreated[ticket_date] = append(index.reversedMapTicketsCreated[ticket_date], index.tickets[i])

		assignee_by_name := index.reversedMapUsersByName[strings.ToLower(assignee_by_id.Name)]
		for j := 0; j < len(assignee_by_name); j++ {
			assignee_by_name[j].Tickets = append(assignee_by_name[j].Tickets, index.tickets[i])
		}
		index.reversedMapUsersByName[strings.ToLower(assignee_by_id.Name)] = assignee_by_name

		name_splice := strings.Split(assignee_by_id.Name, " ")
		for _, name := range name_splice {
			assignee_by_split_name := index.reversedMapUsersByName[strings.ToLower(name)]
			for j := 0; j < len(assignee_by_split_name); j++ {
				if assignee_by_split_name[j].Id == id {
					assignee_by_split_name[j].Tickets = append(assignee_by_split_name[j].Tickets, index.tickets[i])
				}
			}
			index.reversedMapUsersByName[strings.ToLower(name)] = assignee_by_split_name
		}

		assignee_by_flag := index.reversedMapUsersByVerify[strconv.FormatBool(assignee_by_id.Verified)]
		for j := 0; j < len(assignee_by_flag); j++ {
			if assignee_by_flag[j].Id == id {
				assignee_by_flag[j].Tickets = append(assignee_by_flag[j].Tickets, index.tickets[i])
			}
		}
		index.reversedMapUsersByVerify[strings.ToLower(assignee_by_id.Name)] = assignee_by_flag

		assignee_by_date := index.reversedMapUsersCreated[assignee_by_id.CreatedAt.Format("2006-01-02")]
		for j := 0; j < len(assignee_by_date); j++ {
			if assignee_by_date[j].Id == id {
				assignee_by_date[j].Tickets = append(assignee_by_date[j].Tickets, index.tickets[i])
			}
		}
		index.reversedMapUsersCreated[assignee_by_id.CreatedAt.Format("2006-01-02")] = assignee_by_date
	}
}
