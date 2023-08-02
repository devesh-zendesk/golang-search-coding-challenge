package userInterface

import (
	"bufio"
	"fmt"
	"golang-search-coding-challenge/model"
	"os"
	"strings"
)

func UserInterface() {
	userFinder := model.Init()
	for {
		fmt.Println("\nHow Do you want to Search?")
		fmt.Println("Accepted Inputs: id, name, verified, exit")
		var action string
		fmt.Scan(&action)

		processCommand(action, userFinder)
	}
}

func processCommand(action string, userFinder model.UserFinder) {
	reader := bufio.NewReader(os.Stdin)
	switch action {
	case "id":
		fmt.Println("Please enter the ID you want to search")
		var input int
		fmt.Scan(&input)
		user, err := userFinder.GetUserById(input)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nUser Details:\n\tUser ID: ", user.Id, "\n\tName: ", user.Name, "\n\tCreated At: ", user.CreatedAt, "\n\tVerified: ", user.Verified)
			fmt.Println("\nTicket Details:")
			for _, item := range user.Tickets {
				fmt.Println("\tTicket ID: ", item.Id, "\n\tType: ", item.Type, "\n\tSubject: ", item.Subject, "\n\tCreated At: ", item.CreatedAt, "\n\tTags: ", item.Tags)
			}
		}

	case "name":
		fmt.Println("Please enter the Name you want to search")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		user, err := userFinder.GetUserByName(input_string)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, item := range user {
			fmt.Println("\nUser Details:\n\tUser ID: ", item.Id, "\n\tName: ", item.Name, "\n\tCreated At: ", item.CreatedAt, "\n\tVerified: ", item.Verified)
			fmt.Println("\nTicket Details:")
			for _, ele := range item.Tickets {
				fmt.Println("\n\tTicket ID: ", ele.Id, "\n\tType: ", ele.Type, "\n\tSubject: ", ele.Subject, "\n\tCreated At: ", ele.CreatedAt, "\n\tTags: ", ele.Tags)
			}
		}
	case "verified":
		fmt.Println("Please enter the Verify Flag you want to search")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		user, err := userFinder.GetUserByVerifiedFlag(input_string)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, item := range user {
			fmt.Println("\n\tUser ID: ", item.Id, "\n\tName: ", item.Name, "\n\tCreated At: ", item.CreatedAt, "\n\tVerified: ", item.Verified, "\n.")
		}

	case "exit":
		os.Exit(0)

	default:
		fmt.Println("Invalid")
	}
}
