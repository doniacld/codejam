// This CodeJam template was freely built upon this one:
// http://weblog.shank.in/input-template-go-for-algorithmic-competitions/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// //////////////////////////////////////////////////////////////////////////////

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

// //////////////////////////////////////////////////////////////////////////////

func main() {
	mi := MyInput{rdr: os.Stdin}
	run(mi, func(s string) {
		fmt.Println(s)
	})
}

// //////////////////////////////////////////////////////////////////////////////

func run(mi MyInput, out func(string)) {
	t := mi.readInt()
	for caseNo := 1; caseNo <= t; caseNo++ {
		n := mi.readInt()
		counts := make([]int, 4)

		for l := 0; l < n; l++ {
			line := mi.readLine()
			counts[0] += countIs(line[:n])
			counts[1] += countIs(line[n:])
		}

		for l := n; l < 2*n; l++ {
			line := mi.readLine()
			counts[2] += countIs(line[:n])
			counts[3] += countIs(line[n:])
		}

		out(fmt.Sprintf("Case #%d: %d", caseNo, touch(counts)))
	}
}

func countIs(line string) int {
	count := 0
	for _, l := range strings.Split(line, "") {
		if l == "I" {
			count++
		}
	}
	return count
}

func touch(nbsIs []int) int {
	return diff(nbsIs[0], nbsIs[3]) + diff(nbsIs[1], nbsIs[2])
}

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
