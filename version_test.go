package version

import (
	"testing"
)

func TestRunMain(t *testing.T) {
	err := Run()
	if err != nil {
		t.Errorf("Got error: %s", err.Error())
	}
}
