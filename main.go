package main

import (
	"fmt"
	"golang-search-coding-challenge/model"
)

// The data struct for the decoded data
// Notice that all fields must be exportable!

func main() {

	// variable to get user input for ID search
	var idFromUser = 0

	fmt.Println("Please enter the id you want to Search!")
	fmt.Scanln(&idFromUser)
	// fmt.Println(reflect.TypeOf(idFromUser))
	searchUsed(idFromUser)

}

func searchUsed(searchID any) {

	payload := model.LoadUsersJson()
	info := model.LoadTicketJson()

	for i := 0; i < len(payload); i++ {
		if payload[i].Id == searchID {
			fmt.Println("user id", payload[i].Id)
			fmt.Println("user name ", payload[i].Name)
			fmt.Println("created at", payload[i].Created_at)
			fmt.Println("verified", payload[i].Verified)
		}
	}

	for j := 0; j < len(info); j++ {
		if info[j].Assignee_id == searchID {
			fmt.Println("Suject", info[j].Subject)
			fmt.Println("User tags", info[j].Tags)
		}
	}

}
