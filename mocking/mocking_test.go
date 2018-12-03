package main

import (
	"bytes"
	"testing"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Count int
}

func (s *SpySleeper) Sleep() {
	s.Count++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := SpySleeper{}
	Countdown(&buffer, &sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if sleeper.Count != 4 {
		t.Errorf("Countdown didn't call Sleep() 4 times")
	}
}
