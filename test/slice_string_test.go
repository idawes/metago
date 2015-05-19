package test

import (
	"testing"
)

func TestSliceString(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    
    if !a.Equals(&b) {
        t.Errorf("a and b should be equal")
    }

    a.VString = append(a.VString, "Foo")
    b.VString = append(b.VString, "Bar")

    if a.Equals(&b) {
        t.Errorf("a and b should not be equal")
    }

	d := a.Diff(&b)
	if len(d.Chgs) != 1 {
		t.Errorf("Found %d differences between a and b, expected 1", len(d.Chgs))
	}

	a.Apply(d)
	if !a.Equals(&b) {
		t.Errorf("a and b should be equal")
	}
}
