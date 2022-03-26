// http://weblog.shank.in/input-template-go-for-algorithmic-competitions/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	// f, _ := os.Open("./input.txt")
	// mi := MyInput{rdr: f}
	mi := MyInput{rdr: os.Stdin}

	t := mi.readInt()
	for caseNo := 1; caseNo <= t; caseNo++ {
		_ = mi.readLine()
		res := mi.readInts()
		printOutput(caseNo, getMin(res))
	}
}

func getMin(sizes []int) int {
	sort.Ints(sizes)

	treats := 1
	sum := 0
	prec := sizes[0]
	for _, size := range sizes {
		if size != prec {
			treats++
			prec = size
		}
		sum += treats
	}
	return sum
}

// printOutput prints the output in the requested format
// e.g.: "Case #1: 2"
func printOutput(caseNb, out int) {
	fmt.Printf("Case #%d: %d", caseNb, out)
}
