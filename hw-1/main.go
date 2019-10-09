package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	currTime := time.Now()
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal("Error:\n", err)
	}

	fmt.Printf("Current time: %v | Exact time (NTP): %v\n", currTime.Format("15:04:05"), ntpTime.Format("15:04:05"))
}
