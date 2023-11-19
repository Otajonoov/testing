package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep(duration time.Duration)
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return len(p), nil
}

type ConfigurableSleeper struct {
	duration time.Duration
	Sleeper
}

//-------------

type SpyTime struct {
	durationSleep time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSleep = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.Sleeper.Sleep(c.duration)
}

//-------------

func Countdown(out io.Writer, slepper *SpyCountdownOperations) {
	for i := 3; i > 0; i-- {
		slepper.Sleep()
		fmt.Fprintln(out, i)

	}
	fmt.Fprint(out, "Go!")
}
