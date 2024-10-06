package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// Todo struct
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time //pointer cause it can be null
}

// ToDo slice
type Todos []Todo

// Add todo
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

// Validate index
func (todos *Todos) validateIndex(index int) error {
	// Check if index is valid
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

// Delete todo
func (todos *Todos) delete(index int) error {
	t := *todos //copy of todos since we dont't want to modify the original slice

	// Check if index is valid
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	//t[:index] -> all elements before index excluding index
	//t[index+1:] -> all elements after index excluding index
	// append -> append all elements before index and after index to create a new slice
	// ... -> unpack the slice to individual elements
	//so in conclusion, we are skipping the element at index and creating a new slice

	return nil
}

// Toogle todo(complete or incomplete)
func (todos *Todos) toggle(index int) error {
	t := *todos //copy of todos since we dont't want to modify the original slice

	// Check if index is valid
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed //get the current status of the todo

	if !isCompleted {
		completionTime := time.Now()           //get the current time
		t[index].CompletedAt = &completionTime //set the completion time
	}
	t[index].Completed = !isCompleted //toggle the status

	return nil
}

// Edit todo
func (todos *Todos) edit(index int, title string) error {
	t := *todos //copy of todos since we dont't want to modify the original slice

	// Check if index is valid
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title //edit the title

	return nil
}

// Print all todos
func (todos *Todos) print() {
	table := table.New(os.Stdout) //create a new table
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format("time.RFC1123")
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
