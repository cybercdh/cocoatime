package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Apple Core Data reference date: January 1, 2001, 00:00:00 UTC
var coreDataReferenceDate = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

// convertCoreDataTimestamp converts the Core Data timestamp to a human-readable date
func convertCoreDataTimestamp(timestamp float64) time.Time {
	// Add the seconds from the timestamp to the reference date
	return coreDataReferenceDate.Add(time.Duration(timestamp) * time.Second)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cocoatime <timestamp>")
		os.Exit(1)
	}

	// Get the timestamp from command-line argument
	timestampArg := os.Args[1]

	// Convert argument to float64
	timestamp, err := strconv.ParseFloat(timestampArg, 64)
	if err != nil {
		fmt.Println("Invalid timestamp. Please provide a valid Core Data timestamp.")
		os.Exit(1)
	}

	// Convert the Core Data timestamp to a human-readable date
	humanReadableDate := convertCoreDataTimestamp(timestamp)

	// Print both the input timestamp and the human-readable date
	fmt.Printf("%s,%s\n", timestampArg, humanReadableDate.Format(time.RFC3339))

}
