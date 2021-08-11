package controllers

import (
	"github.com/ideaspaper/puttask/services"
	"github.com/ideaspaper/puttask/views"
)

type ITodoController interface {
	Help()
	List(sortFlag *string)
	Add(*string)
	FindById(*int)
	FindByTask(*string)
	Delete(*int)
	Complete(*int)
	Uncomplete(*int)
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

func (tc *todoController) List(sortFlag *string) {
	todos, err := tc.todoService.List(sortFlag)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.List(todos)
	}
}

func (tc *todoController) Add(newTask *string) {
	todo, err := tc.todoService.Add(newTask)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Add(todo)
	}
}

func (tc *todoController) FindById(id *int) {
	todo, err := tc.todoService.FindById(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.FindById(todo)
	}
}

func (tc *todoController) FindByTask(task *string) {
	todos := tc.todoService.FindByTask(task)
	tc.todoView.FindByTask(todos)
}

func (tc *todoController) Delete(id *int) {
	todo, err := tc.todoService.Delete(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Delete(todo)
	}
}

func (tc *todoController) Complete(id *int) {
	todo, err := tc.todoService.Complete(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Complete(todo)
	}
}

func (tc *todoController) Uncomplete(id *int) {
	todo, err := tc.todoService.Uncomplete(id)
	if err != nil {
		tc.todoView.DisplayError(err)
	} else {
		tc.todoView.Uncomplete(todo)
	}
}
