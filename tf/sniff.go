package tf

import (
	"encoding/json"
	"os/exec"
	"strings"
)

type versionJson struct {
	TerraformVersion string `json:"terraform_version"`
}

func versionFromTool(tool string) string {
	if tool == "_" {
		return ""
	}

	cmd := exec.Command(tool, "version", "-json")
	raw, _ := cmd.Output()

	parsed := versionJson{}
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return ""
	}

	return parsed.TerraformVersion + "?"
}

func getVersion(version []string, tool string) string {
	if len(version) > 0 {
		return strings.Join(version, "/")
	}

	return versionFromTool(tool)
}

func Sniff(version ...string) Version {
	tool := getParentCommand()

	return Version{
		Version: getVersion(version, tool),
		Tool:    tool,
	}
}
