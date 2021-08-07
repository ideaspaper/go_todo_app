package controllers

import (
	"github.com/ideaspaper/go_todo_app/models"
	"github.com/ideaspaper/go_todo_app/views"
)

func Help() {
	views.Help()
}

func List() {
	data, err := models.List()
	if err != nil {
		views.DisplayError(err)
	} else {
		views.List(data)
	}
}

func Add(newTask string) {
	newTodo, err := models.Add(newTask)
	if err != nil {
		views.DisplayError(err)
	} else {
		views.Add(newTodo)
	}
}

func FindById(id string) {
	foundTodo, err := models.FindById(id)
	if err != nil {
		views.DisplayError(err)
	} else {
		views.FindById(foundTodo)
	}
}

func Delete(id string) {
	deletedTodo, err := models.Delete(id)
	if err != nil {
		views.DisplayError(err)
	} else {
		views.Delete(deletedTodo)
	}
}

func Complete(id string) {
	completedTodo, err := models.Complete(id)
	if err != nil {
		views.DisplayError(err)
	} else {
		views.Complete(completedTodo)
	}
}

func Uncomplete(id string) {
	uncompletedTodo, err := models.Uncomplete(id)
	if err != nil {
		views.DisplayError(err)
	} else {
		views.Uncomplete(uncompletedTodo)
	}
}
