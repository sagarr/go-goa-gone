package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

// default sleeper with 1 second hardcoded duration
type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// configurable sleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}


func Countdown(sleeper Sleeper, out io.Writer) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, "Go!")
}

func main() {
	// sleeper := &DefaultSleeper{}
	sleeper := &ConfigurableSleeper{5 * time.Second, time.Sleep}
	Countdown(sleeper, os.Stdout)
}
