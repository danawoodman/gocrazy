package gocrazy

import (
	"time"
)

/*
ExecuteAtInterval runs a function at a specified interval, which can be dynamically
updated at runtime by passing a new interval to the intervalChan channel.

Passing an interval of 0 to the intervalChan channel will pause the ticker.

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
	var ticker *time.Ticker
	if interval > 0 {
		ticker = time.NewTicker(interval)
		defer ticker.Stop()
	}

	for {
		if ticker != nil {
			select {
			case <-ticker.C:
				fn()
			case newInterval, ok := <-intervalChan:
				if !ok {
					return // channel closed, exit goroutine
				}
				if newInterval > 0 {
					if ticker != nil {
						ticker.Stop()
					}
					ticker = time.NewTicker(newInterval)
				} else if newInterval == 0 {
					if ticker != nil {
						ticker.Stop()
						ticker = nil // Pause ticker
					}
				}
			}
		} else {
			// Ticker is paused, wait for new interval
			newInterval, ok := <-intervalChan
			if !ok {
				return // channel closed, exit goroutine
			}
			if newInterval > 0 {
				ticker = time.NewTicker(newInterval)
			}
		}
	}
}
