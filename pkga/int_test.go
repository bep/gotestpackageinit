package pkga_test

import (
	"testing"

	"github.com/bep/gotestpackageinit/pkgb"
)

func TestB(t *testing.T) {
	if pkgb.B() != "b" {
		t.Fatal("pkgb.B() != b")
	}
}
