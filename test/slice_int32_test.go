package test

import (
	"testing"
)

func TestSliceInt32(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    
    if !a.Equals(&b) {
        t.Errorf("a and b should be equal")
    }

    a.VInt32 = append(a.VInt32, 3)
    b.VInt32 = append(b.VInt32, 5)

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
