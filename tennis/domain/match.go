package domain

import "sync"

type Match struct {

}

func (m Match) Start(p1 *Player, p2 *Player) {
	wg := &sync.WaitGroup{}
	c := make(chan string)

	wg.Add(1)
	go p1.Play(c, wg)
	c<-"ball"
	wg.Add(1)
	go p2.Play(c, wg)

	wg.Wait()
}
