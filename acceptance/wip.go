package acceptance

import (
	"os"
	"strings"
	"testing"
)

// WIP allows skipping acceptance tests. This should only be used at an acceptance test level as committed unit tests
// should always be in a passing state
func WIP(t *testing.T, ref string) {
	t.Helper()

	flag, ok := os.LookupEnv("WIP_ACCEPTANCE_ENABLED")
	if !ok || !shouldRunTests(ref, flag) {
		t.Skip("skipping WIP acceptance test")
	}
}

func shouldRunTests(wantedRef string, flag string) bool {
	if flag == "true" {
		return true
	}

	references := strings.Split(flag, ",")
	for _, ref := range references {
		if ref == wantedRef {
			return true
		}
	}
	return false
}
