//Exercise: Skipping the Type or Value When Declaring Variables

package main

import (
	"fmt"
	"time"
)

// Declaring Multiple Variables at Once with var without types
var (
	Debug       bool // bool's default value is falsecd
	LogLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
