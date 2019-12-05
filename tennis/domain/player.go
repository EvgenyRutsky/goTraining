package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	Name string
	Skill int //Skill value represents the chance to hit the ball (0-9)
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func CreatePlayer (name string, skill int) (*Player, error){
	if skill > 9 || skill < 0 {
		err := errors.New("wrong skill value has been assigned to Player, expect [0:9] value")
		return nil, err
	}
	newPlayer := &Player{
		Name:  name,
		Skill: skill,
	}
	return newPlayer, nil
}

func (p Player) Play (c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		ball, ok := <- c
		if !ok {
			fmt.Println(p.Name, "has won")
			return
		}
		if rand.Intn(10) > p.Skill {
			fmt.Println(p.Name, "misses")
			close(c)
			return
		}
		fmt.Println(p.Name, "hits")
		c <- ball
	}
}
