package acceptance

import (
	"os"
	"testing"
)

// WIP allows skipping acceptance tests. This should only be used at an acceptance test level as committed unit tests
// should always be in a passing state
func WIP(t *testing.T) {
	t.Helper()

	if shouldRunWIPTests, _ := os.LookupEnv("WIP_ACCEPTANCE_ENABLED"); shouldRunWIPTests != "true" {
		t.Skip("skipping WIP acceptance test")
	}
}
