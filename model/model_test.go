package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var now = time.Now()

func TestLoadIndex(t *testing.T) {
	//Prepare tests
	copy_index := GetRequiredData()

	//Run the tests
	copy_index.LoadIndex()

	//check the results
	assert.True(t, len(copy_index.reversedMapUsersById) > 0)
}

func TestTicketsWithoutAssignee(t *testing.T) {
	//Preparation of the test
	copy_index := GetRequiredData()

	//exicuting the test
	copy_index.LoadIndex()
	actual_ticket, _ := copy_index.GetTicketsWithoutAssignee()

	//check results
	expected_ticket := []Ticket{
		{
			Id:         "asdfasfsaasdf",
			CreatedAt:  now,
			Type:       "problem",
			Subject:    "A Catastrophe in Pune",
			AssigneeId: 0,
			Tags:       []string{"asdf", "asdf", "asdfsdf"},
		},
	}
	assert.Equal(t, expected_ticket, actual_ticket)
}

func GetRequiredData() Index {
	users := []User{
		{
			Id:        1,
			Name:      "Milind Shinde",
			CreatedAt: now,
			Verified:  true,
		},
		{
			Id:        2,
			Name:      "Devesh Chinchole",
			CreatedAt: now,
			Verified:  false,
		},
		{
			Id:        3,
			Name:      "Bruno Marques",
			CreatedAt: now,
			Verified:  true,
		},
	}

	tickets := []Ticket{
		{
			Id:         "asdfasfsa",
			CreatedAt:  now,
			Type:       "incident",
			Subject:    "A Catastrophe in Hungary",
			AssigneeId: 1,
			Tags:       []string{"asdf", "asdf", "asdfsdf"},
		},
		{
			Id:         "asdfasfsaasdf",
			CreatedAt:  now,
			Type:       "task",
			Subject:    "A Catastrophe in Melbourn",
			AssigneeId: 2,
			Tags:       []string{"asdf", "asdf", "asdfsdf"},
		},
		{
			Id:         "asdfasfsaasdf",
			CreatedAt:  now,
			Type:       "problem",
			Subject:    "A Catastrophe in Pune",
			AssigneeId: 0,
			Tags:       []string{"asdf", "asdf", "asdfsdf"},
		},
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

	return index
}
