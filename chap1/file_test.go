package chap1

import (
	"testing"
)

func TestOpenFile(t *testing.T) {
	if err := fileOps(); err != nil {
		t.Fatalf("%+v", err)
	}

}
