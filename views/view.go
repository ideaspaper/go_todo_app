package views

import (
	"fmt"
	"time"

	"github.com/ideaspaper/go_todo_app/models"
)

func Help() {
	fmt.Println("Todo App is a tool for managing your todo list\n\n" +
		"Usage:\n\n" +
		"  todoapp <command> [arguments]\n\n" +
		"The commands are:\n\n" +
		"  help                   display help\n" +
		"  list                   display todo list\n" +
		"  add <task>             add task to todo list\n" +
		"  findById <task_id>     find todo list item by its id\n" +
		"  delete <task_id>       delete todo list item by its id\n" +
		"  complete <task_id>     mark a todo list item as complete\n" +
		"  uncomplete <task_id>   mark a todo list item as uncomplete\n")
}

func List(data []models.Todo) {
	if len(data) == 0 {
		fmt.Printf("Nothing to do, use add command to create a task\n")
	} else {
		for _, v := range data {
			complete := "[ ]"
			if v.Status() {
				complete = "[x]"
			}
			fmt.Printf("%v. %v %v [Added: %v]\n", v.Id(), complete, v.Task(), v.Added().Format(time.RFC822))
		}
	}
}

func Add(newTodo models.Todo) {
	fmt.Printf("%v has been added successfully\n", newTodo.Task())
}

func FindById(foundTodo models.Todo) {
	complete := "[ ]"
	if foundTodo.Status() {
		complete = "[x]"
	}
	fmt.Printf("%v. %v %v [Added: %v]\n", foundTodo.Id(), complete, foundTodo.Task(), foundTodo.Added())
}

func Delete(deletedTodo models.Todo) {
	fmt.Printf("[%v. %v] has been deleted successfully\n", deletedTodo.Id(), deletedTodo.Task())
}

func Complete(completedTodo models.Todo) {
	fmt.Printf("[%v. %v] has been marked as complete\n", completedTodo.Id(), completedTodo.Task())
}

func Uncomplete(completedTodo models.Todo) {
	fmt.Printf("[%v. %v] has been marked as uncomplete\n", completedTodo.Id(), completedTodo.Task())
}

func DisplayError(err error) {
	fmt.Println(err)
}
