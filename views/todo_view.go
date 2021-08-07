package views

import (
	"fmt"
	"time"

	"github.com/ideaspaper/go_todo_app/entities"
)

type ITodoView interface {
	Help()
	List([]entities.Todo)
	Add(entities.Todo)
	FindById(entities.Todo)
	Delete(entities.Todo)
	Complete(entities.Todo)
	Uncomplete(entities.Todo)
	DisplayError(error)
}

type todoView struct{}

func NewTodoView() ITodoView {
	return &todoView{}
}

func (tv *todoView) Help() {
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

func (tv *todoView) List(todos []entities.Todo) {
	if len(todos) == 0 {
		fmt.Printf("Nothing to do, use add command to create a task\n")
	} else {
		for _, v := range todos {
			complete := "[ ]"
			if v.Status() {
				complete = "[x]"
			}
			fmt.Printf("%v. %v %v [Added: %v]\n", v.Id(), complete, v.Task(), v.Added().Format(time.RFC822))
		}
	}
}

func (tv *todoView) Add(newTodo entities.Todo) {
	fmt.Printf("[%v. %v] has been added successfully\n", newTodo.Id(), newTodo.Task())
}

func (tv *todoView) FindById(foundTodo entities.Todo) {
	complete := "[ ]"
	if foundTodo.Status() {
		complete = "[x]"
	}
	fmt.Printf("%v. %v %v [Added: %v]\n", foundTodo.Id(), complete, foundTodo.Task(), foundTodo.Added())
}

func (tv *todoView) Delete(deletedTodo entities.Todo) {
	fmt.Printf("[%v. %v] has been deleted successfully\n", deletedTodo.Id(), deletedTodo.Task())
}

func (tv *todoView) Complete(completedTodo entities.Todo) {
	fmt.Printf("[%v. %v] has been marked as complete\n", completedTodo.Id(), completedTodo.Task())
}

func (tv *todoView) Uncomplete(uncompletedTodo entities.Todo) {
	fmt.Printf("[%v. %v] has been marked as uncomplete\n", uncompletedTodo.Id(), uncompletedTodo.Task())
}

func (tv *todoView) DisplayError(err error) {
	fmt.Println(err)
}
