package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceInt16(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // make a longer than b
    a.VInt16 = append(a.VInt16, 3)
    testSliceInt16DiffAndApply(t, a, b)
    
    // make a and b the same length, but with a change.
    b.VInt16 = append(b.VInt16, 5)
    testSliceInt16DiffAndApply(t, a, b)

    // make a shorter than b
    a = SliceTestObject{}
    testSliceInt16DiffAndApply(t, a, b)

    // make both non-nil, and a longer than b
    a.VInt16 = append(a.VInt16, 5)
    a.VInt16 = append(a.VInt16, 3)
    testSliceInt16DiffAndApply(t, a, b)

    // make both same length, but with a change
    b.VInt16 = append(b.VInt16, 5)
    testSliceInt16DiffAndApply(t, a, b)
    
    // make both non-nil, and a shorter than b
    a.VInt16 = a.VInt16[:len(a.VInt16)-1]
    testSliceInt16DiffAndApply(t, a, b)
}

func testSliceInt16DiffAndApply(t *testing.T, a, b SliceTestObject) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, 1, len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
