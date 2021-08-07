package services

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/ideaspaper/go_todo_app/entities"
)

type ITodoService interface {
	List() ([]entities.Todo, error)
	Add(string) (entities.Todo, error)
	FindById(string) (entities.Todo, error)
	Delete(string) (entities.Todo, error)
	Complete(string) (entities.Todo, error)
	Uncomplete(string) (entities.Todo, error)
	Save() error
}

type todoService struct {
	todos []entities.Todo
}

func NewTodoService(fileName string) (ITodoService, error) {
	v, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	todoJson := []entities.TodoJson{}
	if err = json.Unmarshal(v, &todoJson); err != nil {
		return nil, err
	}
	todos := []entities.Todo{}
	for _, v := range todoJson {
		newTodo, err := entities.NewTodo(v.Id, v.Task, v.Added, v.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, *newTodo)
	}

	return &todoService{
		todos,
	}, nil
}

func (ts *todoService) List() ([]entities.Todo, error) {
	return ts.todos, nil
}

func (ts *todoService) Add(newTask string) (entities.Todo, error) {
	newId := 1
	if len(ts.todos) != 0 {
		newId = ts.todos[len(ts.todos)-1].Id() + 1
	}
	newTodo, _ := entities.NewTodo(newId, newTask, time.Now().Format(time.RFC3339), false)
	ts.todos = append(ts.todos, *newTodo)
	err := ts.Save()
	if err != nil {
		return entities.Todo{}, err
	}
	return *newTodo, nil
}

func findById(id, left int, right int, todos []entities.Todo) (int, error) {
	if right >= left {
		mid := left + (right-left)/2
		if todos[mid].Id() == id {
			return mid, nil
		}
		if todos[mid].Id() > id {
			return findById(id, left, mid-1, todos)
		}
		return findById(id, mid+1, right, todos)
	}
	return -1, errors.New("no record found")
}

func (ts *todoService) FindById(id string) (entities.Todo, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return entities.Todo{}, err
	}
	index, err := findById(idInt, 0, len(ts.todos)-1, ts.todos)
	if err != nil {
		return entities.Todo{}, err
	}
	return ts.todos[index], nil
}

func (ts *todoService) Delete(id string) (entities.Todo, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return entities.Todo{}, err
	}
	index, err := findById(idInt, 0, len(ts.todos)-1, ts.todos)
	if err != nil {
		return entities.Todo{}, err
	}
	deletedTodo := ts.todos[index]
	newTodos := append(ts.todos[:index], ts.todos[index+1:]...)
	ts.todos = newTodos
	err = ts.Save()
	if err != nil {
		return entities.Todo{}, err
	}
	return deletedTodo, nil
}

func (ts *todoService) changeStatus(id string, status bool) (entities.Todo, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return entities.Todo{}, err
	}
	index, err := findById(idInt, 0, len(ts.todos)-1, ts.todos)
	if err != nil {
		return entities.Todo{}, err
	}
	ts.todos[index].SetStatus(status)
	err = ts.Save()
	if err != nil {
		return entities.Todo{}, err
	}
	return ts.todos[index], nil
}

func (ts *todoService) Complete(id string) (entities.Todo, error) {
	completedTodo, err := ts.changeStatus(id, true)
	if err != nil {
		return entities.Todo{}, err
	}
	return completedTodo, nil
}

func (ts *todoService) Uncomplete(id string) (entities.Todo, error) {
	completedTodo, err := ts.changeStatus(id, false)
	if err != nil {
		return entities.Todo{}, err
	}
	return completedTodo, nil
}

func (ts *todoService) Save() error {
	todosJson := []entities.TodoJson{}
	for _, v := range ts.todos {
		newTodoJson, err := entities.NewTodoJson(v.Id(), v.Task(), v.Added().Format(time.RFC3339), v.Status())
		if err != nil {
			return err
		}
		todosJson = append(todosJson, *newTodoJson)
	}
	toWrite, _ := json.MarshalIndent(todosJson, "", "  ")
	os.WriteFile("todo_list.json", toWrite, 0666)
	return nil
}
