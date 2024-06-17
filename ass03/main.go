package main

import (
	"fmt"
	"os"
	"strconv"
)

type Response int64

const (
	TooLow  Response = 0
	TooHigh Response = 1
	Lost    Response = 2
	Won     Response = 3
	Go      Response = 4
)

type Guess struct {
	guesser int
	value   int
}

func main() {

	fmt.Println("Booted.")

	fmt.Println("Reading arguments.")
	argsWithProg := os.Args
	numPlayers, errAtoi := strconv.Atoi(argsWithProg[1])
	if errAtoi != nil || numPlayers < 1 {
		fmt.Println("Wrong argument N for players number")
		panic(errAtoi)
	}
	maxNumber, errAtoi := strconv.Atoi(argsWithProg[2])
	if errAtoi != nil || maxNumber < 1 {
		fmt.Println("Wrong argument MAX_NUMBER for max number to guess")
		panic(errAtoi)
	}

	/* channels used by players to submit guesses */
	toOracleChannel := make(chan Guess)
	/* channels used by the oracle to submit results */
	var toPlayerChannels []chan Response

	/* spawning players */
	for i := 0; i < numPlayers; i++ {
		toPlayerChannel := spawnDumbPlayer(i, toOracleChannel, maxNumber)

		/* collecting to player channels */
		toPlayerChannels = append(toPlayerChannels, toPlayerChannel)
	}

	var endChannel chan int
	go Oracle(toOracleChannel, toPlayerChannels, numPlayers, maxNumber, endChannel)

	<-endChannel

}
