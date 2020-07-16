package chap1

import (
	"testing"
)

func TestServer(t *testing.T) {
	if err := serverDemo(); err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestClient(t *testing.T) {
	if err := clientDemo(); err != nil {
		t.Fatalf("%+v", err)
	}

}
