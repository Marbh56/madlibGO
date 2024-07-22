package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a noun:")
	noun, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get noun.")
		return
	}
	noun = strings.TrimSpace(noun)

	fmt.Print("Enter a verb:")
	verb, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get verb.")
		return
	}
	verb = strings.TrimSpace(verb)

	fmt.Print("Enter an adjective:")
	adjective, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get adjective.")
		return
	}
	adjective = strings.TrimSpace(adjective)

	fmt.Print("Enter a adverb:")
	adverb, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get adverb.")
		return
	}
	adverb = strings.TrimSpace(adverb)

	fmt.Printf("Do you %s your %s %s %s?\n", verb, adjective, noun, adverb)

}
