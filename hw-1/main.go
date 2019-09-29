package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

type clock struct {
	year  int
	month int
	day   int
}

func main() {
	currentcClock := new(clock)
	ntpClock := new(clock)

	currTime := time.Now()
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err == nil {
		currentcClock.year, currentcClock.month, currentcClock.day = currTime.Clock()
		ntpClock.year, ntpClock.month, ntpClock.day = ntpTime.Clock()

		fmt.Printf("Current time: %v:%v:%v | Exact time (NTP): %v:%v:%v", currentcClock.year, currentcClock.month, currentcClock.day, ntpClock.year, ntpClock.month, ntpClock.day)
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
