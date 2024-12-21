package main

import "fmt"

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Reading", "Coding", "Playing"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice1 = slice1[1:]
	fmt.Println(slice1)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Learn Go", "Learn Python"}
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Learn Java"
	goals = append(goals, "Learn C#")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	products := []Product{
		{"A Book", "1234", 10.99},
		{"A Carpet", "5678", 45.99},
	}

	products = append(products, Product{"A Sofa", "3456", 200.0})
	fmt.Println(products)
}

type Product struct {
	title string
	id    string
	price float64
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
