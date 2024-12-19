package main

import (
	"baolam/note/note"
	"baolam/note/todo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

/*type displayer interface {
	Display()
}*/

type outputable interface {
	saver
	Display()
}

/*type outputable interface {
	Save() error
	Display()
}*/

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
	}

	strVal, ok := value.(string)

	if ok {
		fmt.Println("String:", strVal)
	}
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func getTodoData() string {
	return getUserInput("Note content:")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
