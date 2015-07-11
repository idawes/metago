package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceUint8(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // make a longer than b
    a.VUint8 = append(a.VUint8, 3)
    testSliceUint8DiffAndApply(t, a, b, 1)
    
    // make a and b the same length, but with a change.
    b.VUint8 = append(b.VUint8, 5)
    testSliceUint8DiffAndApply(t, a, b, 1)

    // make a shorter than b
    a = SliceTestObject{}
    testSliceUint8DiffAndApply(t, a, b, 1)

    // make both non-nil, and a longer than b
    a.VUint8 = append(a.VUint8, 5)
    a.VUint8 = append(a.VUint8, 3)
    testSliceUint8DiffAndApply(t, a, b, 1)

    // make both same length, but with a change
    b.VUint8 = append(b.VUint8, 5)
    testSliceUint8DiffAndApply(t, a, b, 1)
    
    // make both non-nil, and a shorter than b
    a.VUint8 = a.VUint8[:len(a.VUint8)-1]
    testSliceUint8DiffAndApply(t, a, b, 1)
}

func testSliceUint8DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
