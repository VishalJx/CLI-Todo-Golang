package main

func main() {
	// Create a new todo
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos) // Load the todos from the file
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)

	storage.Save(todos) // Save the todos to the file
}
