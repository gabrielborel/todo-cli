package todo

import (
	"errors"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(id int) error {
	ls := *t
	if id <= 0 || id > len(ls) {
		return errors.New("invalid index")
	}

	ls[id-1].CompletedAt = time.Now()
	ls[id-1].Done = true
	return nil
}

func (t *Todos) Delete(id int) error {
	ls := *t
	if id <= 0 || id > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:id-1], ls[id:]...)
	return nil
}
