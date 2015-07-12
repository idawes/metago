package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceUint32(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VUint32 = append(a.VUint32, 3) // sa = {VA}, sb = nil
    testSliceUint32DiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VUint32 = append(b.VUint32, 5) // sa = {VA}, sb = {VB}
    testSliceUint32DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceUint32DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VUint32 = append(a.VUint32, 5)
    a.VUint32 = append(a.VUint32, 3) // sa = {VB, VA}, sb = {VB}
    testSliceUint32DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VUint32 = append(b.VUint32, 5) // sa = {VB, VA}, sb = {VB, VB}
    testSliceUint32DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VUint32 = a.VUint32[:len(a.VUint32)-1] // sa = {VB}, sb = {VB, VB}
    testSliceUint32DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VUint32 = append(a.VUint32, 3)
    a.VUint32 = append(a.VUint32, 3)
    a.VUint32 = append(a.VUint32, 3)
    a.VUint32 = append(a.VUint32, 3) 
    b = SliceTestObject{}
    b.VUint32 = append(b.VUint32, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testSliceUint32DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VUint32[0] = 3
    b.VUint32 = append(b.VUint32, 5) 
    b.VUint32 = append(b.VUint32, 5) 
    b.VUint32 = append(b.VUint32, 5) 
    b.VUint32 = append(b.VUint32, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testSliceUint32DiffAndApply(t, a, b, 5)
    
}

func testSliceUint32DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
