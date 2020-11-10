package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// ComandLineInterface is responsible for handle the user inputs
func comandLineInterface() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi, with this program we'll find the best route for your travel!")
	fmt.Println("---------------------")

	for {
		fmt.Print("\nplease enter the route: ")
		userInput, _ := reader.ReadString('\n')
		cleanInput := strings.Replace(userInput, "\n", "", -1) // convert CRLF to LF
		input := strings.Split(cleanInput, "-")

		if len(input) != 2 || input[0] == "" || input[1] == "" {
			fmt.Print("Please, enter a input with this format: GRU-CGD\n\n")
			time.Sleep(500 * time.Millisecond)
			continue
		}

		from := input[0]
		to := input[1]

		fmt.Println(from, to)

		time.Sleep(700 * time.Millisecond)
		fmt.Print("\n\nPress x and hit enter to try again or any other key to finish the program: ")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal("Error to process the input: ", err)
		}

		switch char {
		case 'x':
			continue
		default:
			fmt.Println("Thank you!")
			return
		}
	}
}
