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

    // single deletion diff at index 0
    a.VFloat32 = append(a.VFloat32, 3.34) // sa = {VA}, sb = nil
    testSliceFloat32DiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VFloat32 = append(b.VFloat32, 5.42) // sa = {VA}, sb = {VB}
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VFloat32 = append(a.VFloat32, 5.42)
    a.VFloat32 = append(a.VFloat32, 3.34) // sa = {VB, VA}, sb = {VB}
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VFloat32 = append(b.VFloat32, 5.42) // sa = {VB, VA}, sb = {VB, VB}
    testSliceFloat32DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VFloat32 = a.VFloat32[:len(a.VFloat32)-1] // sa = {VB}, sb = {VB, VB}
    testSliceFloat32DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VFloat32 = append(a.VFloat32, 3.34)
    a.VFloat32 = append(a.VFloat32, 3.34)
    a.VFloat32 = append(a.VFloat32, 3.34)
    a.VFloat32 = append(a.VFloat32, 3.34) 
    b = SliceTestObject{}
    b.VFloat32 = append(b.VFloat32, 5.42) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testSliceFloat32DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VFloat32[0] = 3.34
    b.VFloat32 = append(b.VFloat32, 5.42) 
    b.VFloat32 = append(b.VFloat32, 5.42) 
    b.VFloat32 = append(b.VFloat32, 5.42) 
    b.VFloat32 = append(b.VFloat32, 5.42) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testSliceFloat32DiffAndApply(t, a, b, 5)
    
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
