// http://weblog.shank.in/input-template-go-for-algorithmic-competitions/

package main

import (
	"bufio"
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
	f, _ := os.Open("input_file.in")
	mi := MyInput{rdr: f}
	// mi := MyInput{rdr: os.Stdin}

	t := mi.readInt()
	for caseNo := 1; caseNo <= t; caseNo++ {
		// TODO - solve the case !
	}
}
