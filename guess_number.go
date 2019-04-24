package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type GuessingGame struct {
	Guesses int
	Name    string
	Number  int32
}

var suppressGuessMessage bool

const GuessName = "Guesser"
const GuessingChances = 6
const MaxNumber = 100

func main() {

	gg := GuessingGame{0, "", 0}
	gg.Greet()
	gg.Choose_Number()
	gg.Play()
}

func (gg *GuessingGame) Choose_Number() {
	if gg.Number == 0 {
		rand.Seed(time.Now().UnixNano())
		gg.Number = rand.Int31n(MaxNumber) + 1
		fmt.Println(gg.Number)
	}
}

func (gg *GuessingGame) Greet() {
	fmt.Println("Hello! What is your Name?\n ")
	fmt.Scanf("%s", &gg.Name)
	if gg.Name == "" {
		gg.Name = GuessName
	}
	fmt.Printf("Well, %s, I am thinking of a number between 1 and 100.\n", gg.Name)
}

func (gg *GuessingGame) makeGuess() bool {
	var guess int32

	gg.askContestant(&guess)
	if guess > 0 {
		if gg.Guesses == 5 {
			fmt.Printf("You lose %s, the number we needed was %d", gg.Name, gg.Number)
			os.Exit(0)
		}
		if guess < gg.Number {
			fmt.Printf("Your %d is too low.\n", guess)
		}
		if guess > gg.Number {
			fmt.Printf("Your %d is too high.\n", guess)
		}
		if guess == gg.Number {
			fmt.Printf("Good job, %s! You guessed my number in %d guesses!", gg.Name, gg.Guesses)
			return false
		}
	}
	return true
}

func (gg *GuessingGame) askContestant(guess *int32) {
	var s string
	if suppressGuessMessage != true {
		fmt.Println("\nTake a guess !!")
	}
	fmt.Scanln(&s)
	//fmt.Println(s)
	i, _ := strconv.Atoi(s)
	*guess = int32(i)
	if s == "" {
		suppressGuessMessage = true
	} else {
		suppressGuessMessage = false
	}
}

func (gg *GuessingGame) Play() {
	for gg.Guesses <= GuessingChances && gg.makeGuess() {
		gg.Guesses++
	}

	if gg.Guesses == GuessingChances {
		fmt.Printf("Sorry, but you've run out of guesses. The number I had in mind was %d\n", gg.Number)
	}
}
