package test

import (
	"testing"
)

func TestSliceInt(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    
    if !a.Equals(&b) {
        t.Errorf("a and b should be equal")
    }

    a.VInt = append(a.VInt, 3)
    b.VInt = append(b.VInt, 5)

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
