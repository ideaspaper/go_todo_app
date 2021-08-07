package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ideaspaper/go_todo_app/controllers"
	"github.com/ideaspaper/go_todo_app/services"
	"github.com/ideaspaper/go_todo_app/views"
)

func main() {
	fileName := "todo_list.json"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("file %v does not exist\n", fileName)
		fmt.Printf("creating %v\n", fileName)
		os.WriteFile("todo_list.json", []byte("[]"), 0666)
		fmt.Printf("file %v created successfully\n", fileName)
	}

	todoService, err := services.NewTodoService(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	todoController := controllers.NewTodoController(todoService, views.NewTodoView())

	if len(os.Args) < 2 {
		fmt.Println(errors.New("wrong commands/options"))
		todoController.Help()
		os.Exit(1)
	}

	command := os.Args[1]

	if command == "help" {
		todoController.Help()
	} else if command == "list" {
		todoController.List()
	} else {
		if len(os.Args) < 3 {
			fmt.Println(errors.New("wrong commands/options, see help"))
			os.Exit(1)
		}
		switch command {
		case "add":
			options := os.Args[2:]
			todoController.Add(options[0])
		case "findById":
			options := os.Args[2:]
			todoController.FindById(options[0])
		case "delete":
			options := os.Args[2:]
			todoController.Delete(options[0])
		case "complete":
			options := os.Args[2:]
			todoController.Complete(options[0])
		case "uncomplete":
			options := os.Args[2:]
			todoController.Uncomplete(options[0])
		default:
			todoController.Help()
		}
	}
}
