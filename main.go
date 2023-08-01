package main

import (
	"bufio"
	"fmt"
	"golang-search-coding-challenge/model"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	userFinder := model.Init()
	fmt.Println("How Do you want to Search?")
	fmt.Println("Accepted Inputs: id, name")
	var action string
	fmt.Scan(&action)

	switch action {
	case "id":
		fmt.Println("Please enter the ID you want to search")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string, _ := strconv.Atoi(input)
		user, err := userFinder.GetUserById(input_string)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("User ID: ", user.Id, "\nName: ", user.Name, "\nCreated At: ", user.CreatedAt, "\nVerified: ", user.Verified)

	case "name":
		fmt.Println("Please enter the Name you want to search")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input_string := strings.ToLower(input)
		user, err := userFinder.GetUserByName(input_string)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("User ID: ", user.Id, "\nName: ", user.Name, "\nCreated At: ", user.CreatedAt, "\nVerified: ", user.Verified)

	default:
		fmt.Println("Invalid")
	}
}
