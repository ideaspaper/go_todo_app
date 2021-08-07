package controllers

import (
	"github.com/ideaspaper/go_todo_app/services"
	"github.com/ideaspaper/go_todo_app/views"
)

type ITodoController interface {
	Help()
	List()
	Add(string)
	FindById(string)
	Delete(string)
	Complete(string)
	Uncomplete(string)
}

type todoController struct {
	todoService services.ITodoService
	todoView    views.ITodoView
}

func NewTodoController(todoService services.ITodoService, todoView views.ITodoView) ITodoController {
	return &todoController{
		todoService: todoService,
		todoView:    todoView,
	}
}

func (tc *todoController) Help() {
	tc.todoView.Help()
}

func (tc *todoController) List() {
	todos, err := tc.todoService.List()
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.List(todos)
	}
}

func (tc *todoController) Add(newTask string) {
	todo, err := tc.todoService.Add(newTask)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Add(todo)
	}
}

func (tc *todoController) FindById(id string) {
	todo, err := tc.todoService.FindById(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.FindById(todo)
	}
}

func (tc *todoController) Delete(id string) {
	todo, err := tc.todoService.Delete(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Delete(todo)
	}
}

func (tc *todoController) Complete(id string) {
	todo, err := tc.todoService.Complete(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Complete(todo)
	}
}

func (tc *todoController) Uncomplete(id string) {
	todo, err := tc.todoService.Uncomplete(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Uncomplete(todo)
	}
}
