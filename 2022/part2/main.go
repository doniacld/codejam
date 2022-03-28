// This CodeJam template was freely built upon this one:
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

		// line 1: nb Delivery, the number of orders and the amount of leaves needed for each order
		nbDelivery, _, neededLeaves := readLine1(mi.readInts())

		// line 2-nbDelivery: Deliveries
		deliveries := make(map[int]delivery, 0)
		for i := 0; i < nbDelivery; i++ {
			d := readDeliveries(mi.readInts())
			deliveries[d.minute] = d
		}

		// last line:
		orderTimes := mi.readInts()
		possibleOrders := orders(orderTimes, deliveries, neededLeaves)

		out(fmt.Sprintf("Case #%d: %d", caseNo, possibleOrders))
	}
}

func readLine1(values []int) (int, int, int) {
	return values[0], values[1], values[2]
}

type delivery struct {
	minute    int
	leaves    int
	rotenTime int
}

func readDeliveries(line []int) delivery {
	return delivery{
		minute:    line[0],
		leaves:    line[1],
		rotenTime: line[0] + line[2],
	}
}

func orders(orders []int, deliveries map[int]delivery, neededLeaves int) int {
	totalLeaves := computeActualLeaves(deliveries)
	count := 0
	removedLeaves := 0
	for i := range orders {
		// est ce que j'ai assez de feuille ?
		totalLeaves[i] -= removedLeaves
		if neededLeaves <= totalLeaves[i] {
			// possible
			count++
			removedLeaves = neededLeaves
		}
	}

	return count
}

func computeActualLeaves(deliveries map[int]delivery) map[int]int {
	totalLeaves := make(map[int]int, 0)
	for _, d := range deliveries {
		totalLeaves[d.rotenTime] -= d.leaves + deliveries[d.rotenTime].leaves
	}
	return totalLeaves
}
