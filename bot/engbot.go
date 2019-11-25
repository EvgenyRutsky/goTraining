package bot

import (
	"fmt"
	"time"
)

//Engbot struct
type Engbot struct {
	name string
}

//SayHello says hello
func (e *Engbot) SayHello() {
	fmt.Printf("Hello, I'm %v\n", e.name)
}

//SayTime says Time
func (e *Engbot) SayTime() {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Now is", t.Format("15:04"))
}

//SayDate says date
func (e *Engbot) SayDate() {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Today is", t.Format("January 2 2006"))
}

//SayWeekDate says weekday
func (e *Engbot) SayWeekDay() {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Today is", t.Format("Monday"))
}

//SayBye says bye
func (e *Engbot) SayBye() {
	fmt.Println("Bye")
}

//PrintError prints the error when command isn't recognized
func (e *Engbot) PrintError() {
	fmt.Println("I don't know such command")
}
