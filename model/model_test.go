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
		reversedMapUsersById:     make(map[int]User),
		reversedMapUsersByName:   make(map[string][]User),
		reversedMapUsersByVerify: make(map[string][]User),
		users:                    users,
		tickets:                  tickets,
	}

	//Run the tests
	index.LoadIndex()

	//check the results
	assert.True(t, len(index.reversedMapUsersById) > 0)
}
