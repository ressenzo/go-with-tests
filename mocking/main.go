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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const countdownStart = 3
const finalWorld = "Go!"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWorld)
}

func main() {
	sleeper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
