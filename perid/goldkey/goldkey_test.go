package goldkey

import (
	"testing"
)

func TestCastString(t *testing.T) {
	var testdata string
	testdata = CastString("'foobar'")

	if testdata != "foobar" {
		t.Errorf("string includes quote")
	}
	
	testdata = CastString("\"foobar\"")
	if testdata != "foobar" {
		t.Errorf("string includes quote")
	}
}
