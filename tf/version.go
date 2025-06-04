package tf

import (
	"fmt"
	"strings"
)

type Version struct {
	Version string
	Tool    string
}

func (v Version) String() string {
	str := fmt.Sprintf("%s %s", v.Tool, v.Version)
	return strings.TrimSpace(str)
}

func (v Version) UserAgentPostfix() string {
	str := v.String()

	// Len is <= 1 when both Tool and Version are unknown.
	if len(str) > 1 {
		return fmt.Sprintf(" (%s)", str)
	}
	return ""
}
