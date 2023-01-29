package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gabrielborel/todo-cli"
)

const (
	file = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(file); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Any task")
		err := todos.Store(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid command")
		os.Exit(1)
	}
}
