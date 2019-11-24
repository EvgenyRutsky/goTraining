package bot

import (
	"fmt"
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
	fmt.Println("Сейчас ", t.Location(), t.Format("15:04"))
}

//SayDate says date
func (r *Rubot) SayDate() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Сегодня ", t.Location(), t.Format("Jan 2 2006"))
}

//SayWeekDate says weekdate
func (r *Rubot) SayWeekDay() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println("Сегодня ", t.Location(), t.Format("Mon"))
}

//SayBye says bye
func (r *Rubot) SayBye() {
	fmt.Println("Пока")
}
