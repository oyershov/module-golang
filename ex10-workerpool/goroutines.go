package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Wg struct {
	id int
}

func (w *Wg) Run(j float64, jobs chan float64, wg chan *Wg) {
	for {
		fmt.Printf("worker:%d sleep:%.1f\n", w.id, j)
		time.Sleep(time.Duration(float64(j) * float64(time.Second)))
		select {
		case j = <-jobs:
		default:
			fmt.Printf("worker:%d stopping\n", w.id)
			wg <- w
			return
		}
	}
}

func Scheduler(poolSize int, jobs chan float64) {
	wg := make(chan *Wg, poolSize)

	for i := 1; i <= poolSize; i++ {
		wg <- &Wg{i}
	}

	for {
		w := <-wg
		j := <-jobs
		fmt.Printf("worker:%d spawning\n", w.id)
		go w.Run(j, jobs, wg)
	}
}

func Reader(jobs chan float64) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return
		}
		jobs <- f
	}
}

func Run(poolSize int) {
	jobs := make(chan float64)
	go Scheduler(poolSize, jobs)
	go Reader(jobs)
}
