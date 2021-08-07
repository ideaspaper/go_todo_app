package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ideaspaper/go_todo_app/controllers"
)

func main() {
	if _, err := os.Stat("todo_list.json"); os.IsNotExist(err) {
		fmt.Println("file todo_list.json does not exist")
		fmt.Println("creating todo_list.json")
		os.WriteFile("todo_list.json", []byte("[]"), 0666)
		fmt.Println("file todo_list.json created successfully")
	}

	if len(os.Args) < 2 {
		fmt.Println(errors.New("wrong commands/options"))
		controllers.Help()
		os.Exit(1)
	}

	command := os.Args[1]

	if command == "help" {
		controllers.Help()
	} else if command == "list" {
		controllers.List()
	} else {
		if len(os.Args) < 3 {
			fmt.Println(errors.New("wrong commands/options, see help"))
			os.Exit(1)
		}
		switch command {
		case "add":
			options := os.Args[2:]
			controllers.Add(options[0])
		case "findById":
			options := os.Args[2:]
			controllers.FindById(options[0])
		case "delete":
			options := os.Args[2:]
			controllers.Delete(options[0])
		case "complete":
			options := os.Args[2:]
			controllers.Complete(options[0])
		case "uncomplete":
			options := os.Args[2:]
			controllers.Uncomplete(options[0])
		default:
			controllers.Help()
		}
	}
}
