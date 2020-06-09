// Exercise: Declaring Multiple Variables at Once with var

package main

import (
	"fmt"
	"time"
)

// Declaring Multiple Variables at Once with var
var (
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
