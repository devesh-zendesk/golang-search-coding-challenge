package main

import (
	"encoding/json"
	"fmt"
	"golang-search-coding-challenge/model"
	"io/ioutil"
	"log"
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

	// Let's first read the `config.json` and `tickets.json` file
	userscontent, usercontenterr := ioutil.ReadFile("./users.json")
	ticketscontent, ticketcontenterr := ioutil.ReadFile("./tickets.json")

	if usercontenterr != nil {
		log.Fatal("Error when opening file: ", usercontenterr)
	} else if ticketcontenterr != nil {
		log.Fatal("Erro when opening file: ", ticketcontenterr)
	}

	// Now let's unmarshall the data into `payload`
	var payload model.Users
	var info model.UsersTickets

	ticketcontenterr = json.Unmarshal(ticketscontent, &info)
	usercontenterr = json.Unmarshal(userscontent, &payload)
	if usercontenterr != nil {
		log.Fatal("Error during Unmarshal(): ", usercontenterr)
	} else if ticketcontenterr != nil {
		log.Fatal("Error during Unmarshal(): ", ticketcontenterr)
	}

	// for _, val := range payload {
	// 	// Let's print the unmarshalled data!
	// 	if val.Id == searchID {
	// 		fmt.Printf("id: %v\n", val.Id)
	// 		fmt.Printf("name: %v\n", val.Name)
	// 		fmt.Printf("created: %v\n", val.Created_at)
	// 		fmt.Printf("verfied: %v\n", val.Verified)
	// 		fmt.Printf("")
	// 		break
	// 	}
	// }

	for i := 0; i < len(payload); i++ {
		if payload[i].Id == searchID {

			fmt.Println("user id", payload[i].Id)
			fmt.Println("user name ", payload[i].Name)
			fmt.Println("created at", payload[i].Created_at)
			fmt.Println("verified", payload[i].Verified)
			// fmt.Println("UserTickets", payload[i].info)
		}
	}

	for j := 0; j < len(info); j++ {
		if info[j].Assignee_id == searchID {
			fmt.Println("Suject", info[j].Subject)
			fmt.Println("User tags", info[j].Tags)
		}
	}

}
