package main

import "fmt"

// Struct — groups related fields under one type
type Person struct {
	Name  string
	Age   int
	Email string
}

// Value receiver — read-only, works on a copy
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s, age %d.", p.Name, p.Age)
}

// Pointer receiver — can modify the original
func (p *Person) Birthday() {
	p.Age++
}

// String() — fmt.Println calls this automatically (fmt.Stringer interface)
func (p Person) String() string {
	return fmt.Sprintf("Person{%s, %d, %s}", p.Name, p.Age, p.Email)
}

// Constructor function — Go convention instead of a constructor keyword
func NewPerson(name string, age int, email string) Person {
	return Person{Name: name, Age: age, Email: email}
}

// Embedding — Employee gets all Person fields and methods promoted
type Address struct {
	City    string
	Country string
}

type Employee struct {
	Person          // embedded — fields and methods promoted to Employee
	Address Address // named nested struct
	Role    string
}

func (e Employee) Details() string {
	return fmt.Sprintf("%s | %s | %s", e.Name, e.Role, e.Address.City)
}

// Interface — any type with these methods satisfies Shape (no "implements" needed)
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (c Circle) Area() float64         { return 3.14159 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64    { return 2 * 3.14159 * c.Radius }

func printShape(s Shape) {
	fmt.Printf("  %T → area: %.2f  perimeter: %.2f\n", s, s.Area(), s.Perimeter())
}

func main() {

	// Creating structs
	fmt.Println("--- creating structs ---")

	p1 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	p2 := NewPerson("Bob", 25, "bob@example.com")

	var p3 Person // zero value — all fields are zero values
	p3.Name = "Carol"
	p3.Age = 28

	fmt.Println(p1) // calls p1.String() automatically
	fmt.Println(p2)
	fmt.Println(p3)

	// Methods
	fmt.Println("\n--- methods ---")

	fmt.Println(p1.Greet())
	fmt.Println("age before:", p1.Age) // 30
	p1.Birthday()                       // pointer receiver modifies p1
	fmt.Println("age after:", p1.Age)  // 31

	// Pointer to struct
	fmt.Println("\n--- pointer to struct ---")

	pp := &Person{Name: "Dave", Age: 40, Email: "dave@example.com"}
	pp.Birthday()                      // Go auto-dereferences for pointer receivers
	fmt.Println("Dave's age:", pp.Age) // 41

	// Struct comparison — works when all fields are comparable
	fmt.Println("\n--- comparison ---")

	a := Person{Name: "Eve", Age: 22, Email: "eve@example.com"}
	b := Person{Name: "Eve", Age: 22, Email: "eve@example.com"}
	fmt.Println("a == b:", a == b) // true

	// Embedding
	fmt.Println("\n--- embedding ---")

	emp := Employee{
		Person:  Person{Name: "Grace", Age: 32, Email: "grace@example.com"},
		Address: Address{City: "Seoul", Country: "Korea"},
		Role:    "Engineer",
	}
	fmt.Println(emp.Name)          // promoted from Person — same as emp.Person.Name
	fmt.Println(emp.Greet())       // promoted method from Person
	fmt.Println(emp.Details())
	fmt.Println(emp.Address.City)  // nested struct accessed directly

	// Interface
	fmt.Println("\n--- interface ---")

	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 5},
	}
	for _, s := range shapes {
		printShape(s)
	}

	// EXERCISE
	// 1. Create a BankAccount struct with Owner string and Balance float64.
	//    Add Deposit(amount float64) and Withdraw(amount float64) error methods.
	//    Return an error if withdrawing more than the balance.
	// 2. Add a Triangle struct that satisfies the Shape interface.
	//    Add it to the shapes slice — printShape needs no changes.
	// 3. Create a Manager struct that embeds Employee and adds a
	//    Reports []Employee field. Add a method TeamSize() int.
}
