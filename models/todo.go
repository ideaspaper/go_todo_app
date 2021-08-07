package models

import "time"

type Todo struct {
	id     int
	task   string
	added  time.Time
	status bool
}

func (t *Todo) Id() int {
	return t.id
}

func (t *Todo) Task() string {
	return t.task
}

func (t *Todo) SetTask(task string) {
	t.task = task
}

func (t *Todo) Added() time.Time {
	return t.added
}

func (t *Todo) Status() bool {
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
