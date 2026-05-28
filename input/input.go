// Package input
package input

import (
	"fmt"

	"github.com/0xlichi/govenom/output"
)

// GetHost takes hostname or IP input
func GetHost() string {
	var host string

	fmt.Print(output.BlueColor("Enter target Hostname or IP (eg: target.com or 10.10.10.10): "))
	fmt.Scanln(&host)

	return host
}
