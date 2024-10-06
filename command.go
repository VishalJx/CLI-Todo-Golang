package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	add    string
	delete int
	edit   string
	toggle int
	list   bool
	help   bool
}

// NewCmdFlags creates a new CmdFlags
func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.add, "add", "", "Add a new todo")
	flag.IntVar(&cf.delete, "delete", -1, "Delete a todo by index")
	flag.StringVar(&cf.edit, "edit", "", "Edit a todo by index")
	flag.IntVar(&cf.toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.list, "list", false, "List all todos")
	flag.BoolVar(&cf.help, "help", false, "Show help")

	flag.Parse()
	return &cf
}

// Display help message
func (cf *CmdFlags) displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  --add <todo>       Add a new todo")
	fmt.Println("  --delete <index>   Delete a todo by index")
	fmt.Println("  --edit <index>:<new value>   Edit a todo by index")
	fmt.Println("  --toggle <index>   Toggle a todo by index")
	fmt.Println("  --list             List all todos")
	fmt.Println("  --help             Display this help message")

	os.Exit(0)
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.help:
		cf.displayHelp()
	case cf.add != "":
		todos.add(cf.add)
	case cf.delete != -1:
		todos.delete(cf.delete)
	case cf.edit != "":
		parts := strings.SplitN(cf.edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit format")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cf.toggle != -1:
		todos.toggle(cf.toggle)
	case cf.list:
		todos.print()
	default:
		fmt.Println("Invalid command")
		cf.displayHelp()
	}
}
