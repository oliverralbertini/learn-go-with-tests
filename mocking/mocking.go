package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(writer, "%d\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
