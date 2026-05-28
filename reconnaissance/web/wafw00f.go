// Package webvuln
package webvuln

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Wafw00f(host string) error {
	out, err := exec.Command("wafw00f", "-a", "-v", host).Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("wafw00f failed: %v", err)))
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "wafw00f/wafw00f-result", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save wafw00f output: %v", err)))
		return err
	}

	fmt.Println(output.Success("wafw00f done. Output saved  to logs/" + host + "/wafw00f/wafw00f-result.txt"))
	return nil
}
