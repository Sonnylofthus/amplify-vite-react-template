package main

import (
	"fmt"
	"time"
)

// Person struct demonstrates basic struct usage
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Todo struct demonstrates a practical example
type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

// Method on Person struct
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Method with pointer receiver (can modify the struct)
func (p *Person) Birthday() {
	p.Age++
}

// Method on Todo struct
func (t Todo) Display() string {
	status := "❌"
	if t.Completed {
		status = "✅"
	}
	return fmt.Sprintf("[%d] %s %s", t.ID, status, t.Title)
}

// TodoList demonstrates working with slices of structs
type TodoList struct {
	Items []Todo
}

// Method to add a todo
func (tl *TodoList) Add(title string) {
	newTodo := Todo{
		ID:        len(tl.Items) + 1,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	tl.Items = append(tl.Items, newTodo)
}

// Method to mark todo as complete
func (tl *TodoList) Complete(id int) {
	for i := range tl.Items {
		if tl.Items[i].ID == id {
			tl.Items[i].Completed = true
			return
		}
	}
}

// Method to display all todos
func (tl TodoList) DisplayAll() {
	fmt.Println("\n=== Todo List ===")
	for _, todo := range tl.Items {
		fmt.Println(todo.Display())
	}
}

func main() {
	fmt.Println("=== Working with Structs ===")

	// Creating a Person
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}

	fmt.Println("\nPerson Details:")
	fmt.Printf("Name: %s\n", person.FullName())
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)

	// Using pointer method to modify struct
	fmt.Println("\nHaving a birthday...")
	person.Birthday()
	fmt.Printf("New age: %d\n", person.Age)

	// Working with TodoList
	todoList := TodoList{}

	// Adding todos
	todoList.Add("Learn Go basics")
	todoList.Add("Build a web server")
	todoList.Add("Create a REST API")
	todoList.Add("Deploy to production")

	// Display all todos
	todoList.DisplayAll()

	// Complete some todos
	fmt.Println("\nCompleting some tasks...")
	todoList.Complete(1)
	todoList.Complete(2)

	// Display updated list
	todoList.DisplayAll()

	// Anonymous struct (useful for one-off data structures)
	fmt.Println("\n=== Anonymous Struct ===")
	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  false,
	}
	fmt.Printf("Server Config: %+v\n", config)
}
