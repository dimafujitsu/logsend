package logsend

import (
	"testing"
)

func TestKube(t *testing.T) {
	KubernetesReader()
	if 1.5 == 1.5 {
		t.Error("Expected 1.5, got ")
	}
}
