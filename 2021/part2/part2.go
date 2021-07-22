// http://weblog.shank.in/input-template-go-for-algorithmic-competitions/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// INPUT TEMPLATE START

type MyInput struct {
	rdr         io.Reader
	lineChan    chan string
	initialized bool
}

func (mi *MyInput) start(done chan struct{}) {
	r := bufio.NewReader(mi.rdr)
	defer func() { close(mi.lineChan) }()
	for {
		line, err := r.ReadString('\n')
		if !mi.initialized {
			mi.initialized = true
			done <- struct{}{}
		}
		mi.lineChan <- strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func (mi *MyInput) readLine() string {
	// if this is the first call, initialize
	if !mi.initialized {
		mi.lineChan = make(chan string)
		done := make(chan struct{})
		go mi.start(done)
		<-done
	}

	res, ok := <-mi.lineChan
	if !ok {
		panic("trying to read from a closed channel")
	}
	return res
}

func (mi *MyInput) readInt() int {
	line := mi.readLine()
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *MyInput) readInt64() int64 {
	line := mi.readLine()
	i, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *MyInput) readInts() []int {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int{}
	for _, s := range parts {
		tmp, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *MyInput) readInt64s() []int64 {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int64{}
	for _, s := range parts {
		tmp, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *MyInput) readWords() []string {
	line := mi.readLine()
	return strings.Split(line, " ")
}

// INPUT TEMPLATE END

////////////////////////////////////////////////////////////////////////////////

func main() {
	// f, _ := os.Open("input.txt")
	// mi := MyInput{rdr: f}
	mi := MyInput{rdr: os.Stdin}

	t := mi.readInt()
	for caseNo := 1; caseNo <= t; caseNo++ {
		_ = mi.readLine()
		res := getLetters(mi.readInts())
		printOutput(caseNo, res)
	}
}

const (
	Achar byte = 'A'
)

func getLetters(input []int) string {
	var letters = "A"

	j := 0
	for j+1 < len(input) {
		nbOdd := input[j]
		nbEven := input[j+1]
		if nbOdd < nbEven {
			letters += climb(nbOdd-1) + unclimb(nbEven+1)
		} else {
			letters += climb(nbOdd) + unclimb(nbEven)
		}
		j += 2
	}

	if len(input)%2 == 1 {
	letters += climb(input[j])
	}

	return letters
}

func climb(nbL int) string {
	letter := Achar + 1
	letters := make([]byte, 0)
	for i := 0; i < nbL; i++ {
		letters = append(letters, letter)
		letter = letter + 1
	}
	return string(letters[:])
}

func unclimb(nbL int) string {
	letter := Achar + byte(nbL)
	letters := make([]byte, 0)
	for i := 0; i < nbL; i++ {
		letter = letter - 1
		letters = append(letters, letter)
	}

	return string(letters[:])
}

// printOutput prints the output in the requested format
// e.g.: "Case #1: 2"
func printOutput(caseNb int, out string) {
	fmt.Printf("Case #%d: %s\n", caseNb, out)
}
