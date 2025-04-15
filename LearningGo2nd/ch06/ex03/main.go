package main

type Person struct {
	Name string
	Age  int
}

func main() {
	people := make([]Person, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		people = append(people, Person{Name: "John", Age: i})
	}
}
