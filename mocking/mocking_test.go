package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type countdownOperationsSpy struct {
	Calls []string
}

func (s *countdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *countdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &countdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!
`
	
		if got != want {
			t.Errorf("got: %q; want: %q", got, want)
		}
	})


	t.Run("Sleep before every print", func(t *testing.T) {
		spySleepPrinter := &countdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("Wanted calls %v; got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("got: %v; want: %v", sleepTime, spyTime.durationSlept)
	}
}
