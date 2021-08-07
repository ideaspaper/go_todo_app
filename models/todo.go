package models

import "time"

type Todo struct {
	id     int
	task   string
	added  time.Time
	status bool
}

func (t *Todo) GetId() int {
	return t.id
}

func (t *Todo) GetTask() string {
	return t.task
}

func (t *Todo) SetTask(task string) {
	t.task = task
}

func (t *Todo) GetAdded() time.Time {
	return t.added
}

func (t *Todo) GetStatus() bool {
	return t.status
}

func (t *Todo) SetStatus(status bool) {
	t.status = status
}

func NewTodo(id int, task string, added string, status bool) (*Todo, error) {
	addedTime, err := time.Parse(time.RFC3339, added)
	if err != nil {
		return nil, err
	}
	return &Todo{
		id:     id,
		task:   task,
		added:  addedTime,
		status: status,
	}, nil
}
