package main

import (
	"golang-search-coding-challenge/model"
	ui "golang-search-coding-challenge/userInterface"
)

func main() {
	userFinder := model.Init()
	ui.UserInterface(userFinder)
}
