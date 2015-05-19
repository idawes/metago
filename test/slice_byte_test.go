package test

import (
	"os"
	"testing"
)

func TestSliceByte(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}

	if !a.Equals(&b) {
		t.Errorf("a and b should be equal")
	}

	a.VByte = append(a.VByte, 3)
	b.VByte = append(b.VByte, 5)

	if a.Equals(&b) {
		t.Errorf("a and b should not be equal")
	}

	d := a.Diff(&b)
	if len(d.Chgs) != 1 {
		t.Errorf("Found %d differences between a and b, expected 1", len(d.Chgs))
		d.WriteIndented(os.Stdout, 0)
	}

	a.Apply(d)
	if !a.Equals(&b) {
		t.Errorf("a and b should be equal")
	}
}
