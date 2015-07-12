package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceFloat32(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // make a longer than b
    a.VFloat32 = append(a.VFloat32, 3.34)
    testSliceFloat32DiffAndApply(t, a, b, 1)
    
    // make a and b the same length, but with a change.
    b.VFloat32 = append(b.VFloat32, 5.42)
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // make a shorter than b
    a = SliceTestObject{}
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // make both non-nil, and a longer than b
    a.VFloat32 = append(a.VFloat32, 5.42)
    a.VFloat32 = append(a.VFloat32, 3.34)
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // make both same length, but with a change
    b.VFloat32 = append(b.VFloat32, 5.42)
    testSliceFloat32DiffAndApply(t, a, b, 1)
    
    // make both non-nil, and a shorter than b
    a.VFloat32 = a.VFloat32[:len(a.VFloat32)-1]
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // make 2 changes
    a.VFloat32[0] = 3.34
    testSliceFloat32DiffAndApply(t, a, b, 2)
}

func testSliceFloat32DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
