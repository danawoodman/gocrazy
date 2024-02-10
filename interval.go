package gocrazy

import (
	"time"
)

/*
ExecuteAtInterval runs a function at a specified interval, which can be dynamically
updated at runtime by passing a new interval to the intervalChan channel.

Usage:

	interval := 1 * time.Second
	intervalChan := make(chan time.Duration)
	go gocrazy.ExecuteAtInterval(func() {
		fmt.Println("Executing function at interval...")
	}, interval, intervalChan)

	// Update interval after some time
	intervalChan <- 2 * time.Second
*/
func ExecuteAtInterval(fn func(), interval time.Duration, intervalChan <-chan time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fn()
		case newInterval := <-intervalChan:
			ticker.Stop()
			ticker = time.NewTicker(newInterval)
		}
	}
}
