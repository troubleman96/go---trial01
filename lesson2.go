package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func greet(name string) string {
	return "Hello, " + name
}

func main() {
	// Variables
	city := "Dar es Salaam"
	age := 20

	// Struct
	u := User{Name: "D Knob", Age: 21}

	// Map
	info := map[string]string{
		"app":  "PangaSafari",
		"lang": "Go",
	}

	// Output
	fmt.Println(city, age)
	fmt.Println(greet("Juma"))
	fmt.Println(u.Name, u.Age)
	fmt.Println(info["app"])
}
