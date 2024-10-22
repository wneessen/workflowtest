package workflowtest

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	if os.Getenv("SKIP_TEST") == "true" {
		t.Skip("Skipping test in CI")
	}
	if Add(1, 2) != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", Add(1, 2))
	}
}
