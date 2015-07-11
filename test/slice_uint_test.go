package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceUint(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // make a longer than b
    a.VUint = append(a.VUint, 3)
    testSliceUintDiffAndApply(t, a, b)
    
    // make a and b the same length, but with a change.
    b.VUint = append(b.VUint, 5)
    testSliceUintDiffAndApply(t, a, b)

    // make a shorter than b
    a = SliceTestObject{}
    testSliceUintDiffAndApply(t, a, b)

    // make both non-nil, and a longer than b
    a.VUint = append(a.VUint, 5)
    a.VUint = append(a.VUint, 3)
    testSliceUintDiffAndApply(t, a, b)

    // make both same length, but with a change
    b.VUint = append(b.VUint, 5)
    testSliceUintDiffAndApply(t, a, b)
    
    // make both non-nil, and a shorter than b
    a.VUint = a.VUint[:len(a.VUint)-1]
    testSliceUintDiffAndApply(t, a, b)
}

func testSliceUintDiffAndApply(t *testing.T, a, b SliceTestObject) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, 1, len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
