//go:build !windows

package tf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getParentCommand() string {
	ppid := fmt.Sprintf("%d", os.Getppid())
	cmd := exec.Command("ps", "-p", ppid, "-o", "comm=")
	raw, _ := cmd.Output()

	tool := strings.TrimSpace(string(raw[:]))
	if tool != "tofu" && tool != "terraform" {
		tool = "_"
	}

	return tool
}
