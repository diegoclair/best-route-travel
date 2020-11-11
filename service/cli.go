package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/diegoclair/best-route-travel/domain/contract"
)

type commandLine struct {
	svc *Service
}

func newCommandLineService(svc *Service) contract.CommandLineService {
	return &commandLine{
		svc: svc,
	}
}

// ComandLineInterface is responsible for handle the user inputs
func (cli *commandLine) RunCLI() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi, with this program we'll find the best route for your travel!")
	fmt.Println("---------------------")

	for {
		fmt.Print("\nplease enter the route: ")
		userInput, _ := reader.ReadString('\n')

		input := cli.cleanInput(userInput)

		inputIsValid := cli.validateInput(input)
		if !inputIsValid {
			continue
		}

		from, to := cli.getRoutes(input)

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

func (cli *commandLine) printUsage() {
	fmt.Print("Please, enter a input with this format: GRU-CGD\n\n")
}

func (cli *commandLine) cleanInput(input string) string {
	return strings.Replace(input, "\n", "", -1) // convert CRLF to LF
}

func (cli *commandLine) validateInput(input string) bool {
	splitedInput := strings.Split(input, "-")
	if len(splitedInput) != 2 || splitedInput[0] == "" || splitedInput[1] == "" {
		cli.printUsage()
		time.Sleep(500 * time.Millisecond)
		return false
	}
	return true
}

func (cli *commandLine) getRoutes(input string) (from, to string) {
	routes := strings.Split(input, "-")
	from, to = routes[0], routes[1]
	return
}
