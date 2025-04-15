package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Complited   bool
	CreatedAt   time.Time
	ComplitedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Complited:   false,
		ComplitedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)

}

func (todos *Todos) validdateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalide index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validdateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validdateIndex(index); err != nil {
		return err
	}

	isComplited := t[index].Complited

	if !isComplited {
		complitedTime := time.Now()
		t[index].ComplitedAt = &complitedTime
	}

	t[index].Complited = !isComplited

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validdateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Название", "Выполнено", "Дата создания", "Дата выполнения")
	for index, t := range *todos {
		comlited := "❌"
		complitedAt := ""

		if t.Complited {
			comlited = "✅"
			if t.ComplitedAt != nil {
				complitedAt = t.ComplitedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, comlited, t.CreatedAt.Format(time.RFC1123), complitedAt)
	}

	table.Render()
}
