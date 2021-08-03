package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("prints the right output", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := &CountdownSpy{}

		Countdown(&buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!
`

		if got != want {
			t.Errorf("Wanted %q, but got %q", want, got)
		}
	})

	t.Run("sleeps before printing", func(t *testing.T) {
		countdownSpy := &CountdownSpy{}

		Countdown(countdownSpy, countdownSpy)

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
		got := countdownSpy.Calls

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted following invocations %v, but got %v", want, got)
		}

	})
}

func TestConfigurationSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	c := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	c.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

const write = "write"
const sleep = "sleep"

type CountdownSpy struct {
	Calls []string
}

func (s *CountdownSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept += duration
}
