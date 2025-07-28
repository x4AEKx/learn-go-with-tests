package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("got %v, but got %v", spySleeperPrinter.Calls, want)
		}
	})

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("got %d, but want 3 calls", spySleeper.Calls)
	}
}

func TestConfigurationSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := &ConfigurationSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("want %v, but got %v", sleepTime, spyTime.durationSlept)
	}
}
