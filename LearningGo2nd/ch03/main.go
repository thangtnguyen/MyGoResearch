package main

import "fmt"

func main() {
	ex01()
	ex02()
	ex03()
}

func ex01() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstTwoValues := greetings[:2]
	fromSecondToFourthValues := greetings[1:4]
	fourthAndFifthValues := greetings[3:5]
	fmt.Println(firstTwoValues)           // Output: [Hello Hola]
	fmt.Println(fromSecondToFourthValues) // Output: [Hola नमस्कार こんにちは]
	fmt.Println(fourthAndFifthValues)     // Output: [こんにちは Привіт]
}

func ex02() {
	var message string = "Hi  and "
	rs := []rune(message)
	fmt.Print(rs[3])
}

func ex03() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	employee1 := Employee{id: 123}
	employee2 := Employee{firstName: "John", lastName: "Doe", id: 456}
	var employee3 Employee
	employee3.firstName = "Jane"
	employee3.lastName = "Smith"
	employee3.id = 789

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
