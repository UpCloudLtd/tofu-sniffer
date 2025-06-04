package tf

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func assertStringsEqual(t *testing.T, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func assertStringMatches(t *testing.T, expected, actual string) {
	t.Helper()

	re, err := regexp.Compile(expected)
	if err != nil {
		t.Fatalf("Invalid regex '%s': %v", expected, err)
	}

	if !re.MatchString(actual) {
		t.Errorf("Expected to match '%s', got '%s'", expected, actual)
	}
}

func runId(val any) string {
	id := fmt.Sprintf("%+v", val)
	id = strings.Trim(id, "{}")
	id = strings.ReplaceAll(id, " ", ",")

	return id
}

func TestSniff_returns_underscore_on_unknown_ppid(t *testing.T) {
	version := Sniff()

	assertStringsEqual(t, "_", version.Tool)
	assertStringsEqual(t, "", version.Version)
}

func TestSniff_joins_versions_with_slashes(t *testing.T) {
	version := Sniff("1.9.1", "windows", "amd64")

	assertStringsEqual(t, "1.9.1/windows/amd64", version.Version)
	assertStringsEqual(t, " (_ 1.9.1/windows/amd64)", version.UserAgentPostfix())
}

func TestSniff_getVersion(t *testing.T) {
	testdata := []struct {
		version  []string
		tool     string
		expected string
	}{
		{
			version:  nil,
			tool:     "tofu",
			expected: `^[0-9]+\.[0-9]+\.[0-9]+\?$`,
		},
		{
			version:  []string{"1.9.1"},
			tool:     "tofu",
			expected: `^1.9.1$`,
		},
		{
			version:  nil,
			tool:     "_",
			expected: "^$",
		},
		{
			version:  nil,
			tool:     "invalid",
			expected: "^$",
		},
	}

	for _, td := range testdata {
		t.Run(runId(td), func(t *testing.T) {
			version := getVersion(td.version, td.tool)
			assertStringMatches(t, td.expected, version)
		})
	}
}
