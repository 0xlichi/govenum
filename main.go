package main

import (
	"fmt"
	"os"

	"github.com/0xlichi/govenom/input"
	"github.com/0xlichi/govenom/ping"
)

func main() {
	// Get hostname/IP from input package
	host := input.GetHost()

	// Check rechability
	if !ping.CheckHost(host) {
		fmt.Printf("Host '%v' is not rechable\n", host)
		os.Exit(1)
	}

	fmt.Printf("Host '%v' is rechable\n", host)
}
