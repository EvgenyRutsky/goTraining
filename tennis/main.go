package main

import (
	"fmt"
	"os"
	"tennis/domain"
)

func main() {
	p1, err := domain.CreatePlayer("Johnson", 5)
	errorCheck(err)
	p2, err := domain.CreatePlayer("Tompson", 5)
	errorCheck(err)

	newMatch := &domain.Match{}
	newMatch.Start(p1, p2)
}

func errorCheck(err error){
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
}
