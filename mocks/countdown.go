package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second, time.Sleep})
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

const start = 3
const end = "Go!"

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := start; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(out, end)

}
