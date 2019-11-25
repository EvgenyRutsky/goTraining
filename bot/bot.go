package bot

import (
	"errors"
	"fmt"
	"os"
)

//Bot interface
type Bot interface {
	SayHello()
	SayDate()
	SayTime()
	SayWeekDay()
	SayBye()
	PrintError()
}

func ScanLanguage() string {
	var language string
	fmt.Println("Please enter your language: English/Russian")
	_, err := fmt.Scanln(&language)

	if err != nil {
		fmt.Printf("You've just ran into error %v\n", err)
		os.Exit(1)
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
		os.Exit(1)
	}

	//input = translateCommand(input)

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

	default:
		b.PrintError()
		HandleInput(b)
	}
}

//translateCommand translates language-specific command to numeric command
//func translateCommand (input string) string {
//	commandMap := map[string]func(){
//		"Привет" : "1",
//		"Время" : "2",
//		"Дата" : "3",
//		"День" : "4",
//		"Пока" : "5",
//		"Hello" : "1",
//		"Time" : "2",
//		"Date" : "3",
//		"Day" : "4",
//		"Bye" : "5",
//	}
//	return commandMap[input]
//}
