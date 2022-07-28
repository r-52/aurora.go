package util

import (
	"testing"
)

func Test_PathHelperGetPath(t *testing.T) {
	path := GetCurrentExecPath()
	if len(path) == 0 {
		t.Error("expected path, got %", path)
	}
}
