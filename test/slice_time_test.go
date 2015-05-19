package test

import (
	"testing"
	"time"
)

func TestSliceTime(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}

	if !a.Equals(&b) {
		t.Errorf("a and b should be equal")
	}

	a.VTime = append(a.VTime, time.Now())
	b.VTime = append(b.VTime, time.Now().Add(5*time.Hour))

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
