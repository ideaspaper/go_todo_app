package models

type TodoJson struct {
	Id     int
	Task   string
	Added  string
	Status bool
}

func NewTodoJson(id int, task string, added string, status bool) (*TodoJson, error) {
	return &TodoJson{
		Id:     id,
		Task:   task,
		Added:  added,
		Status: status,
	}, nil
}
