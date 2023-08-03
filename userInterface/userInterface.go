package userInterface

import (
	"bufio"
	"fmt"
	"golang-search-coding-challenge/model"
	"os"
	"strings"
)

func UserInterface(userFinder model.UserFinder) {
	for {
		fmt.Println("\nSearch By: \n\t1. UserId \n\t2. User Name \n\t3. User Verified Flag \n\t4. User Created Date \n\t5. Type of the Ticket \n\t6. Tag of the Ticket \n\t7. Ticket Created Date \n\t8. Tickets Without Assignee \n\t9. exit")
		var action int
		fmt.Scan(&action)

		processCommand(action, userFinder)
	}
}

func processCommand(action int, userFinder model.UserFinder) {
	reader := bufio.NewReader(os.Stdin)
	switch action {
	case 1:
		fmt.Println("Please enter the ID you want to search")
		var input int
		fmt.Scan(&input)
		user, err := userFinder.GetUserById(input)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nUser Details:\n\tUser ID: ", user.Id, "\n\tName: ", user.Name, "\n\tCreated At: ", user.CreatedAt, "\n\tVerified: ", user.Verified)
			if len(user.Tickets) != 0 {
				fmt.Println("\nTicket Details:")
				for _, item := range user.Tickets {
					fmt.Println("\tTicket ID: ", item.Id, "\n\tType: ", item.Type, "\n\tSubject: ", item.Subject, "\n\tCreated At: ", item.CreatedAt, "\n\tTags: ", item.Tags)
					fmt.Println("")
				}
			}
		}

	case 2:
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
			if len(item.Tickets) != 0 {
				fmt.Println("\nTicket Details:")
				for _, ele := range item.Tickets {
					fmt.Println("\tTicket ID: ", ele.Id, "\n\tType: ", ele.Type, "\n\tSubject: ", ele.Subject, "\n\tCreated At: ", ele.CreatedAt, "\n\tTags: ", ele.Tags)
					fmt.Println("")
				}
			}
		}
	case 3:
		fmt.Println("Please enter the Verify Flag you want to search")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		user, err := userFinder.GetUserByVerifiedFlag(input_string)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, item := range user {
			fmt.Println("\nUser Details:\n\tUser ID: ", item.Id, "\n\tName: ", item.Name, "\n\tCreated At: ", item.CreatedAt, "\n\tVerified: ", item.Verified)
			if len(item.Tickets) != 0 {
				fmt.Println("\nTicket Details:")
				for _, ele := range item.Tickets {
					fmt.Println("\tTicket ID: ", ele.Id, "\n\tType: ", ele.Type, "\n\tSubject: ", ele.Subject, "\n\tCreated At: ", ele.CreatedAt, "\n\tTags: ", ele.Tags)
					fmt.Println("")
				}
			}
		}

	case 4:
		fmt.Println("Please enter the Created Date you want to search. Formate should be YYYY-MM-DD")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		user, err := userFinder.GetUserByCreatedDate(input_string)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, item := range user {
			fmt.Println("\nUser Details:\n\tUser ID: ", item.Id, "\n\tName: ", item.Name, "\n\tCreated At: ", item.CreatedAt, "\n\tVerified: ", item.Verified)
			if len(item.Tickets) != 0 {
				fmt.Println("\nTicket Details:")
				for _, ele := range item.Tickets {
					fmt.Println("\tTicket ID: ", ele.Id, "\n\tType: ", ele.Type, "\n\tSubject: ", ele.Subject, "\n\tCreated At: ", ele.CreatedAt, "\n\tTags: ", ele.Tags)
					fmt.Println("")
				}
			}
		}

	case 5:
		fmt.Println("Please enter the Ticket Type you want to search")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		tickets, err := userFinder.GetTicketsByType(input_string)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, ticket := range tickets {
				fmt.Println("\nUser Details:\n\tUser ID: ", ticket.Assignee.Id, "\n\tName: ", ticket.Assignee.Name, "\n\tCreated At: ", ticket.Assignee.CreatedAt, "\n\tVerified: ", ticket.Assignee.Verified)
				fmt.Println("\nTicket Details:\n\tTicket ID: ", ticket.Id, "\n\tType: ", ticket.Type, "\n\tSubject: ", ticket.Subject, "\n\tCreated At: ", ticket.CreatedAt, "\n\tTags: ", ticket.Tags)
			}
		}

	case 6:
		fmt.Println("Please enter the Tag of the Ticket you want to search")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		tickets, err := userFinder.GetTicketsByTag(input_string)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, ticket := range tickets {
				fmt.Println("\nUser Details:\n\tUser ID: ", ticket.Assignee.Id, "\n\tName: ", ticket.Assignee.Name, "\n\tCreated At: ", ticket.Assignee.CreatedAt, "\n\tVerified: ", ticket.Assignee.Verified)
				fmt.Println("\nTicket Details:\n\tTicket ID: ", ticket.Id, "\n\tType: ", ticket.Type, "\n\tSubject: ", ticket.Subject, "\n\tCreated At: ", ticket.CreatedAt, "\n\tTags: ", ticket.Tags)
			}
		}

	case 7:
		fmt.Println("Please enter the Created Date you want to search. Formate should be YYYY-MM-DD")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		tickets, err := userFinder.GetTicketsByCreatedDate(input_string)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, ticket := range tickets {
				fmt.Println("\nUser Details:\n\tUser ID: ", ticket.Assignee.Id, "\n\tName: ", ticket.Assignee.Name, "\n\tCreated At: ", ticket.Assignee.CreatedAt, "\n\tVerified: ", ticket.Assignee.Verified)
				fmt.Println("\nTicket Details:\n\tTicket ID: ", ticket.Id, "\n\tType: ", ticket.Type, "\n\tSubject: ", ticket.Subject, "\n\tCreated At: ", ticket.CreatedAt, "\n\tTags: ", ticket.Tags)
			}
		}

	case 8:
		tickets, err := userFinder.GetTicketsWithoutAssignee()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, ticket := range tickets {
				fmt.Println("\nTicket Details:\n\tTicket ID: ", ticket.Id, "\n\tType: ", ticket.Type, "\n\tSubject: ", ticket.Subject, "\n\tCreated At: ", ticket.CreatedAt, "\n\tTags: ", ticket.Tags)
			}
		}

	case 9:
		fmt.Println("Terminating the Program")
		os.Exit(0)

	default:
		fmt.Println("Invalid")
	}
}
