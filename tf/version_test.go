package tf

import (
	"testing"
)

func TestVersion(t *testing.T) {
	testdata := []struct {
		tool      string
		version   string
		str       string
		uaPostfix string
	}{
		{
			tool:      "tofu",
			version:   "1.9.1",
			str:       "tofu 1.9.1",
			uaPostfix: " (tofu 1.9.1)",
		},
		{
			tool:      "terraform",
			version:   "1.8.2",
			str:       "terraform 1.8.2",
			uaPostfix: " (terraform 1.8.2)",
		},
	}

	for _, td := range testdata {
		t.Run(runId(td), func(t *testing.T) {
			v := Version{
				Tool:    td.tool,
				Version: td.version,
			}

			assertStringsEqual(t, td.str, v.String())
			assertStringsEqual(t, td.uaPostfix, v.UserAgentPostfix())
		})
	}
}
