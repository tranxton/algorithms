package leet_code

import (
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	firstBadVersion := FirstBadVersion(1000)
	expectedFirstBadVersion := 556

	if firstBadVersion != expectedFirstBadVersion {
		t.Errorf("expected %d, got %d", expectedFirstBadVersion, firstBadVersion)
	}
}
