package main

import (
	"fmt"
	"math/rand"
)

func Player(guessChannel chan Guess, responseChannel chan Response, id int, maxNumber int) {
	lowerLimit := 0
	upperLimit := maxNumber
	over := false
	// wait the Go signal from the Oracle
	<-responseChannel
	for !over {
		// calculate guess
		guessValue := lowerLimit + rand.Intn(upperLimit-lowerLimit)
		guess := Guess{guesser: id, value: guessValue}
		log(id, "Guessing", guessValue)
		//send guess
		guessChannel <- guess
		//wait response
		response := <-responseChannel
		if response == Won {
			log(id, "Yey, I won! ")
		} else if response == TooLow {
			lowerLimit = guessValue
		} else if response == TooHigh {
			upperLimit = guessValue
		}
		over = response == Won || response == Lost
	}
}

func spawnDumbPlayer(id int, guessChannel chan Guess, maxNumber int) chan Response {
	ch := make(chan Response)
	go Player(guessChannel, ch, id, maxNumber)
	return ch
}

func log(id int, a ...any) {
	fmt.Println(" [player-", id, "] ", a)
}
