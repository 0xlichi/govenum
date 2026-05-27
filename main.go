package main

import (
	"fmt"
	"os"

	"github.com/0xlichi/govenom/input"
	"github.com/0xlichi/govenom/output"
	"github.com/0xlichi/govenom/ping"
	"github.com/0xlichi/govenom/ui/banner"
)

func init() {
	banner.GetBanner()
}

func main() {
	// Get hostname/IP from input package
	host := input.GetHost()

	// Check reachability
	if !ping.CheckHost(host) {
		fmt.Println(
			output.Error(fmt.Sprintf("Host '%v' is not reachable.", host)),
		)
		os.Exit(1)
	}

	fmt.Println(
		output.Success(fmt.Sprintf("Host '%v' is reachable.", host)),
	)
}
