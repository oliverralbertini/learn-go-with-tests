package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("slept for '%v' should have slept for '%v'", spyTime.durationSlept, sleepTime)
	}
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Countdown(&buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' wanted '%s'", got, want)
		}
	})

	t.Run("sleeps before each write", func(t *testing.T) {
		spySleepPrinter := CountdownOperationsSpy{}
		Countdown(&spySleepPrinter, &spySleepPrinter)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("Wanted calls '%v' got '%v'", want, spySleepPrinter.Calls)
		}
	})
}

const (
	sleep = "sleep"
	write = "write"
)
