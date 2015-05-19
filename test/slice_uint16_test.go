package test

import (
	"testing"
)

func TestSliceUint16(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    
    if !a.Equals(&b) {
        t.Errorf("a and b should be equal")
    }

    a.VUint16 = append(a.VUint16, 3)
    b.VUint16 = append(b.VUint16, 5)

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
