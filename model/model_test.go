package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadIndex(t *testing.T) {
	//Prepare tests

	users := []User{
		{
			Id:        1,
			Name:      "Milind Shinde",
			CreatedAt: time.Now(),
			Verified:  true,
		},
		{
			Id:        2,
			Name:      "Devesh Chinchole",
			CreatedAt: time.Now(),
			Verified:  true,
		},
	}

	tickets := []Ticket{
		{
			Id:         "asdfasfsa",
			CreatedAt:  time.Now(),
			Type:       "ASDF",
			Subject:    "A Catastrophe in Hungary",
			AssigneeId: 1,
			Tags:       []string{"asdf", "asdf", "asdfsdf"},
		},
		{
			Id:         "asdfasfsaasdf",
			CreatedAt:  time.Now(),
			Type:       "ASDFasasf",
			Subject:    "A Catastrophe in Pune",
			AssigneeId: 2,
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

	//Run the tests
	index.LoadIndex()

	//check the results
	assert.True(t, len(index.reversedMapUsersById) > 0)
}

func TestTicketsWithoutAssignee(t *testing.T) {
	//Preparation of the test
	users := []User{
		{
			Id:        1,
			Name:      "Milind Shinde",
			CreatedAt: time.Now(),
			Verified:  true,
		},
		{
			Id:        2,
			Name:      "Devesh Chinchole",
			CreatedAt: time.Now(),
			Verified:  true,
		},
	}

	tickets := []Ticket{
		{
			Id:         "asdfasfsa",
			CreatedAt:  time.Now(),
			Type:       "incident",
			Subject:    "A Catastrophe in Hungary",
			AssigneeId: 1,
			Tags:       []string{"asdf", "asdf", "asdfsdf"},
		},
		{
			Id:         "asdfasfsaasdf",
			CreatedAt:  time.Now(),
			Type:       "task",
			Subject:    "A Catastrophe in Pune",
			AssigneeId: 2,
			Tags:       []string{"asdf", "asdf", "asdfsdf"},
		},
		{
			Id:         "asdfasfsaasdf",
			CreatedAt:  time.Now(),
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

	//exicuting the test
	index.LoadIndex()
	actual_ticket, _ := index.GetTicketsWithoutAssignee()

	// expected_ticket := []Ticket{
	// 	{
	// 		Id:         "asdfasfsaasdf",
	// 		CreatedAt:  time.Now(),
	// 		Type:       "problem",
	// 		Subject:    "A Catastrophe in Pune",
	// 		AssigneeId: 0,
	// 		Tags:       []string{"asdf", "asdf", "asdfsdf"},
	// 	},
	// }

	//check results
	// assert.Equal(t, expected_ticket, actual_ticket)
	assert.True(t, len(actual_ticket) == 1)
}
