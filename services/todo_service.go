package services

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ideaspaper/puttask/entities"
)

type ITodoService interface {
	List(sortFlag *string) ([]entities.Todo, error)
	Add(*string) (entities.Todo, error)
	FindById(*int) (entities.Todo, error)
	FindByTask(*string) []entities.Todo
	Delete(*int) (entities.Todo, error)
	Complete(*int) (entities.Todo, error)
	Uncomplete(*int) (entities.Todo, error)
	Save() error
}

type todoService struct {
	fileName string
	todos    []entities.Todo
}

func NewTodoService(fileName string) (ITodoService, error) {
	v, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	todosJson := []entities.TodoJson{}
	if err = json.Unmarshal(v, &todosJson); err != nil {
		return nil, err
	}
	todos := []entities.Todo{}
	for _, v := range todosJson {
		newTodo, err := entities.NewTodo(v.Id, v.Task, v.Added, v.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, *newTodo)
	}
	return &todoService{
		fileName,
		todos,
	}, nil
}

func (ts *todoService) List(sortFlag *string) ([]entities.Todo, error) {
	if sortFlag != nil {
		switch *sortFlag {
		case "asc":
			sort.Slice(ts.todos, func(i, j int) bool {
				return strings.ToLower(ts.todos[i].Task()) < strings.ToLower(ts.todos[j].Task())
			})
		case "desc":
			sort.Slice(ts.todos, func(i, j int) bool {
				return strings.ToLower(ts.todos[i].Task()) > strings.ToLower(ts.todos[j].Task())
			})
		default:
			return nil, errors.New("wrong flag value")
		}
	} else {
		sort.Slice(ts.todos, func(i, j int) bool {
			return ts.todos[i].Id() > ts.todos[j].Id()
		})
		sort.Slice(ts.todos, func(i, _ int) bool {
			return !ts.todos[i].Status()
		})
	}
	return ts.todos, nil
}

func (ts *todoService) Add(newTask *string) (entities.Todo, error) {
	newId := 1
	if len(ts.todos) != 0 {
		newId = ts.todos[len(ts.todos)-1].Id() + 1
	}
	newTodo, _ := entities.NewTodo(newId, *newTask, time.Now().Format(time.RFC3339), false)
	ts.todos = append(ts.todos, *newTodo)
	err := ts.Save()
	if err != nil {
		return entities.Todo{}, err
	}
	return *newTodo, nil
}

func findById(id *int, left int, right int, todos []entities.Todo) (int, error) {
	if right >= left {
		mid := left + (right-left)/2
		if todos[mid].Id() == *id {
			return mid, nil
		}
		if todos[mid].Id() > *id {
			return findById(id, left, mid-1, todos)
		}
		return findById(id, mid+1, right, todos)
	}
	return -1, errors.New("no record found")
}

func (ts *todoService) FindById(id *int) (entities.Todo, error) {
	index, err := findById(id, 0, len(ts.todos)-1, ts.todos)
	if err != nil {
		return entities.Todo{}, err
	}
	return ts.todos[index], nil
}

func (ts *todoService) FindByTask(task *string) []entities.Todo {
	result := []entities.Todo{}
	for _, v := range ts.todos {
		if strings.Contains(v.Task(), *task) {
			result = append(result, v)
		}
	}
	return result
}

func (ts *todoService) Delete(id *int) (entities.Todo, error) {
	index, err := findById(id, 0, len(ts.todos)-1, ts.todos)
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

func (ts *todoService) changeStatus(id *int, status bool) (entities.Todo, error) {
	index, err := findById(id, 0, len(ts.todos)-1, ts.todos)
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

func (ts *todoService) Complete(id *int) (entities.Todo, error) {
	completedTodo, err := ts.changeStatus(id, true)
	if err != nil {
		return entities.Todo{}, err
	}
	return completedTodo, nil
}

func (ts *todoService) Uncomplete(id *int) (entities.Todo, error) {
	completedTodo, err := ts.changeStatus(id, false)
	if err != nil {
		return entities.Todo{}, err
	}
	return completedTodo, nil
}

func (ts *todoService) Save() error {
	todosJson := []entities.TodoJson{}
	for _, v := range ts.todos {
		newTodoJson := entities.NewTodoJson(v.Id(), v.Task(), v.Added().Format(time.RFC3339), v.Status())
		todosJson = append(todosJson, *newTodoJson)
	}
	toWrite, _ := json.MarshalIndent(todosJson, "", "  ")
	os.WriteFile(ts.fileName, toWrite, 0666)
	return nil
}
