package service

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/diegoclair/best-route-travel/domain/contract"
)

type commandLine struct {
	svc           *Service
	travelService contract.TravelService
}

func newCommandLineService(svc *Service, travelService contract.TravelService) contract.CommandLineService {
	return &commandLine{
		svc:           svc,
		travelService: travelService,
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

		whereFrom, whereTo := cli.getRoutes(input)
		bestRoute, err := cli.travelService.GetBestRoute(strings.ToUpper(whereFrom), strings.ToUpper(whereTo))
		if err != nil {
			fmt.Print("Error to get best route, contact the admin\n")
			return
		}
		fmt.Printf("best route: %s > $%v", bestRoute.Route, bestRoute.Price)
	}
}

func (cli *commandLine) InputNewFile(inputFileName string) {

	sourceFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("possible_routes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//write a empty line at the end of the file to avoid format problems when use the POST to add new routes
	writer := csv.NewWriter(newFile)
	defer writer.Flush()
	writer.Write([]string{})

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes with success!", bytesCopied)
	return
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
