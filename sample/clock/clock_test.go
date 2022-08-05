package clock_test

import (
	"testing"
	"time"

	"github.com/sanggi-wjg/go_demos/sample/clock"
)

func TestPrintDigit(t *testing.T) {
	clock.PrintDigit()
	/*
	   ███
	   █ █
	   █ █
	   █ █
	   ███

	   ██
	    █
	    █
	    █
	   ███

	   ███
	     █
	   ███
	   █
	   ███

	   ███
	     █
	   ███
	     █
	   ███

	   █ █
	   █ █
	   ███
	     █
	     █

	   ███
	   █
	   ███
	     █
	   ███

	   ███
	   █
	   ███
	   █ █
	   ███

	   ███
	     █
	     █
	     █
	     █

	   ███
	   █ █
	   ███
	   █ █
	   ███

	   ███
	   █ █
	   ███
	     █
	   ███
	*/
}

func TestPrintTime(t *testing.T) {
	now := time.Now()

	hour := now.UTC().Hour()
	minute := now.UTC().Minute()
	second := now.UTC().Second()

	clock.PrintTime(hour, minute, second)
	/*
	   ███  ███       █ █  ███       ███  █ █
	   █ █  █     ░   █ █    █   ░     █  █ █
	   █ █  ███       ███  ███       ███  ███
	   █ █  █ █   ░     █  █     ░     █    █
	   ███  ███         █  ███       ███    █
	*/
}
