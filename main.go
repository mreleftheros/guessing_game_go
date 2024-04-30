package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const MaxGuesses = 10

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func main() {
	luckyNumber, totalGuesses := rand.Intn(100), 1

	for {
		if totalGuesses > MaxGuesses {
			fmt.Println("You lost")
			break
		}

		fmt.Printf("Guesses %d/%d\n", totalGuesses, MaxGuesses)
		fmt.Print("Enter a guess(0-100): ")
		input, err := getInput()
		if err != nil {
			log.Fatal(err)
		}
		guess, err := strconv.ParseInt(input, 10, 32)
		if guess < 0 || guess > 100 || err != nil {
			continue
		}

		fmt.Println("You guessed", guess)
		totalGuesses++
		if guess > int64(luckyNumber) {
			fmt.Println("You guessed too high!")
		} else if guess < int64(luckyNumber) {
			fmt.Println("You guessed too low!")
		} else {
			fmt.Println("Congrats! Lucky number was", luckyNumber)
			break
		}
	}
}
