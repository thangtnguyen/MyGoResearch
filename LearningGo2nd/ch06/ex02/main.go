package main

import "fmt"

func UpdateSlice(slice []string, value string) {
	slice[len(slice)-1] = value
	fmt.Print(slice)
}

func GrowSlice(slice []string, value string) {
	slice = append(slice, value)
	fmt.Print(slice)
}

func main() {
	valSlice := []string{"a", "b", "c"}
	fmt.Print("Before calling UpdateSlice: ")
	fmt.Print(valSlice)
	UpdateSlice(valSlice, "d")
	fmt.Println()
	fmt.Print("After calling UpdateSlice: ")
	fmt.Print(valSlice)
	fmt.Println()
	fmt.Print("Before calling GrowSlice: ")
	fmt.Print(valSlice)
	GrowSlice(valSlice, "e")
	fmt.Println()
	fmt.Print("After calling GrowSlice: ")
	fmt.Print(valSlice)
}
