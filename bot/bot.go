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
		fmt.Printf("You've just ran into error: %v\n", err)
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
		fmt.Printf("You've just ran into error: %v\n", err)
		os.Exit(1)
	}

	commandMap := map[string]func(){
		"Привет" : b.SayHello,
		"Время" : b.SayTime,
		"Дата" : b.SayDate,
		"День" : b.SayWeekDay,
		"Пока" : b.SayBye,
		"Hello" : b.SayHello,
		"Time" : b.SayTime,
		"Date" : b.SayDate,
		"Day" : b.SayWeekDay,
		"Bye" : b.SayBye,
	}
	i, ok := commandMap[input]

	if ok {
		i()
	} else {
		b.PrintError()
	}
	HandleInput(b)
}


