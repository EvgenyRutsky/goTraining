package bot

import (
	"fmt"
	"os"
	"time"
)

//Rubot struct
type Rubot struct {
	name string
}

//SayHello says hello
func (r *Rubot) SayHello() {
	fmt.Printf("Привет, я %v\n", r.name)
}

//SayTime says Time
func (r *Rubot) SayTime() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Сейчас", t.Format("15:04"))
}

//SayDate says date
func (r *Rubot) SayDate() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Сегодня", t.Format("January 2 2006"))
}

//SayWeekDate says weekday
func (r *Rubot) SayWeekDay() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Сегодня", t.Format("Monday"))
}

//SayBye says bye
func (r *Rubot) SayBye() {
	fmt.Println("Пока")
	os.Exit(0)
}

//PrintError prints the error when command isn't recognized
func (r *Rubot) PrintError() {
	fmt.Println("Я не понимаю эту команду")
}
