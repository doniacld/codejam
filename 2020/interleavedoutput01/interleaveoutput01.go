package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read from stdin
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Unable to read from stdin")
	}
	line = strings.Replace(line, "\n", "", -1)

	// retrieve the tests cases number
	tcCounter, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("%s is not a integer", line)
	}

	// solve the test case and print the response
	for i := 0; i < tcCounter; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Unable to read from stdin")
		}
		occ := getOccurrence(line)
		printOutput(i, occ)
	}
}

// getOccurrence returns the number of times the the pattern IO appears
func getOccurrence(in string) int {
	var cI, ci int
	var cOcc int
	events := strings.Split(in, "")
	for _, char := range events {
		switch char {
		case "I":
			cI++
		case "i":
			ci++
		case "O":
			if cI > 0 {
				cI--
				cOcc++
			} else {
				ci--
			}
		case "o":
			if ci > 0 {
				ci--
			} else {
				cI--
			}
		}
	}
	return cOcc
}

// printOutput prints the output in the requested format
// e.g.: "Case #1: 2"
func printOutput(caseNb, out int) {
	msg :=  fmt.Sprintf("Case #%d: %d", caseNb, out)
	fmt.Println(msg)
}
