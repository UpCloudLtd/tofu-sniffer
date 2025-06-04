//go:build windows

package tf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getParentCommand() string {
	cmd := exec.Command("tasklist", "/nh", "/fi", fmt.Sprintf("PID eq %d", os.Getppid()), "/fo", "csv")
	raw, _ := cmd.Output()
	cols := strings.Split(strings.TrimSpace(string(raw[:])), ",")

	tool := strings.TrimSuffix(strings.Trim(cols[0], `"`), ".exe")
	if tool != "tofu" && tool != "terraform" {
		tool = "_"
	}

	return tool
}
