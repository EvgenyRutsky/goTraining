package bot

import (
	"fmt"
)

//Bot interface
type Bot interface {
	createBot()
	sayHello()
	sayDate()
	sayTime()
	sayWeekDay()
	sayBye()
}

//CreateBot creates new bot depending on the language
func CreateBot() {
	var language string
	fmt.Print("Language: English/Russian")
	_, err := fmt.Scanln(&language)

	if err != nil {
		fmt.Printf("You've just ran into error %v\n", err)
	}

	if language == "English" {
		CreateEngbot()
	} else if language == "Russian" {
		CreateRubot()
	} else {
		fmt.Println("Entered language is not valid, try again, please")
		CreateBot()
	}
}

//HandleInput serves an input from the console
func HandleInput(b Bot) {
	var input string
	fmt.Print("You: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Printf("You've just ran into error %v\n", err)
	}

	switch input {
	case "1":
		b.sayHello()
		HandleInput(b)
	case "2":
		b.sayTime()
		HandleInput(b)
	case "3":
		b.sayDate()
		HandleInput(b)
	case "4":
		b.sayWeekDay()
		HandleInput(b)
	case "5":
		b.sayBye()
	}
}
