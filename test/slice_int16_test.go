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

    // single deletion diff at index 0
    a.VInt16 = append(a.VInt16, 3) // sa = {VA}, sb = nil
    testSliceInt16DiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VInt16 = append(b.VInt16, 5) // sa = {VA}, sb = {VB}
    testSliceInt16DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceInt16DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VInt16 = append(a.VInt16, 5)
    a.VInt16 = append(a.VInt16, 3) // sa = {VB, VA}, sb = {VB}
    testSliceInt16DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VInt16 = append(b.VInt16, 5) // sa = {VB, VA}, sb = {VB, VB}
    testSliceInt16DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VInt16 = a.VInt16[:len(a.VInt16)-1] // sa = {VB}, sb = {VB, VB}
    testSliceInt16DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VInt16 = append(a.VInt16, 3)
    a.VInt16 = append(a.VInt16, 3)
    a.VInt16 = append(a.VInt16, 3)
    a.VInt16 = append(a.VInt16, 3) 
    b = SliceTestObject{}
    b.VInt16 = append(b.VInt16, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testSliceInt16DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VInt16[0] = 3
    b.VInt16 = append(b.VInt16, 5) 
    b.VInt16 = append(b.VInt16, 5) 
    b.VInt16 = append(b.VInt16, 5) 
    b.VInt16 = append(b.VInt16, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testSliceInt16DiffAndApply(t, a, b, 5)
    
}

func testSliceInt16DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
