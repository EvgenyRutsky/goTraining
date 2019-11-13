package bot

import (
	"fmt"
	"time"
)

//Engbot struct
type Engbot struct {
	name string
}

//CreateEngbot creates new english bot
func CreateEngbot() *Engbot {
	return &Engbot{
		name: "Bob",
	}
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
	fmt.Println("Now is", t.Location(), t.Format("15:04"))
}

//SayDate says date
func (e *Engbot) SayDate() {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Today is", t.Location(), t.Format("Jan 2 2006"))
}

//SayWeekDate says weekdate
func (e *Engbot) SayWeekDate() {
	name := "Europe/London"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Today is", t.Location(), t.Format("Mon"))
}

//SayBye says bye
func (e *Engbot) SayBye() {
	fmt.Println("Bye")
}
