package bot

import (
	"errors"
	"fmt"
)

//Bot interface
type Bot interface {
	SayHello()
	SayDate()
	SayTime()
	SayWeekDay()
	SayBye()
}

func ScanLanguage() string{
	var language string
	fmt.Println("Language: English/Russian")
	_, err := fmt.Scanln(&language)

	if err != nil {
		fmt.Printf("You've just ran into error %v\n", err)
	}
	return language
}
//CreateBot creates new bot depending on the language
func CreateBot(language string) (Bot, error) {

	if language == "English" {
		var b Bot = &Engbot{
			name: "Bob",
		}
		return b, nil
	} else if language == "Russian" {
		var b Bot = &Rubot{
			name: "Василий",
		}
		return b, nil
	}

	return nil, errors.New("incorrect language has been used")

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
		b.SayHello()
		HandleInput(b)
	case "2":
		b.SayTime()
		HandleInput(b)
	case "3":
		b.SayDate()
		HandleInput(b)
	case "4":
		b.SayWeekDay()
		HandleInput(b)
	case "5":
		b.SayBye()
	}
}
