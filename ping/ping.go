// Package ping
package ping

import "os/exec"

// CheckHost checks if a host is reachable
func CheckHost(host string) bool {
	cmd := exec.Command("ping", "-c", "2", host)

	err := cmd.Run()

	return err == nil
}
