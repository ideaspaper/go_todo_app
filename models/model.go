package models

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
)

func findById(id, left int, right int, todos []Todo) (int, error) {
	if right >= left {
		mid := left + (right-left)/2
		if todos[mid].id == id {
			return mid, nil
		}
		if todos[mid].id > id {
			return findById(id, left, mid-1, todos)
		}
		return findById(id, mid+1, right, todos)
	}
	return -1, errors.New("no record found")
}

func writeToFile(todo []Todo) error {
	todosJson := []TodoJson{}
	for _, v := range todo {
		newTodoJson, err := NewTodoJson(v.id, v.task, v.added.Format(time.RFC3339), v.status)
		if err != nil {
			return err
		}
		todosJson = append(todosJson, *newTodoJson)
	}
	toWrite, _ := json.MarshalIndent(todosJson, "", "  ")
	os.WriteFile("todo_list.json", toWrite, 0666)
	return nil
}

func changeStatus(id string, status bool) (Todo, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Todo{}, err
	}
	var todos []Todo
	todos, err = List()
	if err != nil {
		return Todo{}, err
	}
	index, err := findById(idInt, 0, len(todos)-1, todos)
	if err != nil {
		return Todo{}, err
	}
	todos[index].status = status
	err = writeToFile(todos)
	if err != nil {
		return Todo{}, nil
	}
	return todos[index], nil
}

func List() ([]Todo, error) {
	v, err := os.ReadFile("todo_list.json")
	if err != nil {
		return nil, err
	}
	todoJson := []TodoJson{}
	if err = json.Unmarshal(v, &todoJson); err != nil {
		return nil, err
	}
	todos := []Todo{}
	for _, v := range todoJson {
		var newTodo *Todo
		newTodo, err = NewTodo(v.Id, v.Task, v.Added, v.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, *newTodo)
	}
	return todos, nil
}

func Add(newTask string) (Todo, error) {
	todos, err := List()
	if err != nil {
		return Todo{}, err
	}
	newId := 1
	if len(todos) != 0 {
		newId = todos[len(todos)-1].id + 1
	}
	newTodo, _ := NewTodo(newId, newTask, time.Now().Format(time.RFC3339), false)
	todos = append(todos, *newTodo)
	err = writeToFile(todos)
	if err != nil {
		return Todo{}, err
	}
	return *newTodo, nil
}

func FindById(id string) (Todo, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Todo{}, err
	}
	var todos []Todo
	todos, err = List()
	if err != nil {
		return Todo{}, err
	}
	index, err := findById(idInt, 0, len(todos)-1, todos)
	if err != nil {
		return Todo{}, err
	}
	return todos[index], nil
}

func Delete(id string) (Todo, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Todo{}, err
	}
	var todos []Todo
	todos, err = List()
	if err != nil {
		return Todo{}, err
	}
	index, err := findById(idInt, 0, len(todos)-1, todos)
	if err != nil {
		return Todo{}, err
	}
	newTodos := []Todo{}
	newTodos = append(newTodos, todos[0:index]...)
	newTodos = append(newTodos, todos[index+1:]...)
	err = writeToFile(newTodos)
	if err != nil {
		return Todo{}, nil
	}
	return todos[index], nil
}

func Complete(id string) (Todo, error) {
	completedTodo, err := changeStatus(id, true)
	if err != nil {
		return Todo{}, nil
	}
	return completedTodo, nil
}

func Uncomplete(id string) (Todo, error) {
	uncompletedTodo, err := changeStatus(id, false)
	if err != nil {
		return Todo{}, nil
	}
	return uncompletedTodo, nil
}
