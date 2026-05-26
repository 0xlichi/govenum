// Package input
package input

import "fmt"

// GetHost takes hostname or IP input
func GetHost() string {
	var host string

	fmt.Print("Enter hostname or IP: ")
	fmt.Scanln(&host)

	return host
}
