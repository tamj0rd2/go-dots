package acceptance

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

// example command: WIP_ACCEPTANCE_ENABLED=RIP-123 make acceptance-test
func TestSomeFeature(t *testing.T) {
	WIP(t, "RIP-123")
	t.Log("Running RIP-123")
	assert.Equal(t, 1, 23)
}

// example command: WIP_ACCEPTANCE_ENABLED=RIP-567 make acceptance-test
func TestMyOtherFeature(t *testing.T) {
	WIP(t, "RIP-567")
	t.Log("running RIP-567")
	assert.Equal(t, 5, 67)
}
