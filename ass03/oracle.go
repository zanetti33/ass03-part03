package main

import (
	"fmt"
	"math/rand"
)

func Oracle(guessChannel chan Guess, responseChannels []chan Response, numPlayers int, maxNumber int) {
	magicNumber := rand.Intn(maxNumber)
	someoneGuessedCorrectly := false
	var winningPlayer int
	// send start message
	for i := 0; i < numPlayers; i++ {
		responseChannels[i] <- Go
	}
	for !someoneGuessedCorrectly {
		var guesses []Guess
		oracleLog("Reading all answers")
		// read all answers
		for i := 0; i < numPlayers; i++ {
			guess := <-guessChannel
			oracleLog("Player ", guess.guesser, " guessed the number ", guess.value)
			// someone guesses, if he is the first he wins, other that guessed correctly will still lose
			if !someoneGuessedCorrectly && guess.value == magicNumber {
				oracleLog("Player ", guess.guesser, " guessed the magic number!")
				someoneGuessedCorrectly = true
				winningPlayer = guess.guesser
			}
			guesses = append(guesses, guess)
		}
		oracleLog("Sending response to each player")
		if someoneGuessedCorrectly {
			// if someone guessed correctly he wins, others lose
			for i := 0; i < numPlayers; i++ {
				if i == winningPlayer {
					responseChannels[i] <- Won
				} else {
					responseChannels[i] <- Lost
				}
			}
		} else {
			// otherwise we give a suggestion to eachone
			for i := 0; i < numPlayers; i++ {
				guess := guesses[i]
				if guess.value < magicNumber {
					responseChannels[guess.guesser] <- TooLow
				} else {
					responseChannels[guess.guesser] <- TooHigh
				}
			}
		}
	}
}

func oracleLog(a ...any) {
	fmt.Println(" [Oracle] ", a)
}
