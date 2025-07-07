package projects

import (
	"fmt"
	u "go-learn/utils"
)

func Guess_Machine() {
	const MAX_ALLOWED_TRIES int = 3
	guess := u.Random_Gen(1, 10)
	try_counter := 0

	for {
		var inp int
		fmt.Print("Write your generic guess: ")
		_, err := fmt.Scanln(&inp)

		if err != nil {
			fmt.Printf("Retry > Invalid input, Must be a Number: %v !\n", err)
			continue
		}

		try_counter += 1

		if check_guess(inp, &guess) {
			switch try_counter {
			case 1:
				fmt.Printf("> Huhu Wow 🎉! Congrats, you got it really quick #1 🦁!\n")
			case 2:
				fmt.Printf("> Wow 🤩! Congrats, you got it quite less time #2 🙂!\n")
			case 3:
				fmt.Printf("> Applauses 👏! Congrats, at least you tried #3 😇!\n")
			}
			break
		}

		fmt.Printf("Hmm, it's not right! (Tries: %d/%d)\n", try_counter, MAX_ALLOWED_TRIES)

		// Crossed the Limit:
		if try_counter >= MAX_ALLOWED_TRIES {
			fmt.Printf("> Your guess limit is reached 🥲 !\nIt was: %d\n", guess)
			break
		}
	}
}

func check_guess(n int, guess *int) bool {
	return n == *guess
}
