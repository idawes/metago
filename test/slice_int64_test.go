package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceInt64(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VInt64 = append(a.VInt64, 3) // sa = {VA}, sb = nil
    testSliceInt64DiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VInt64 = append(b.VInt64, 5) // sa = {VA}, sb = {VB}
    testSliceInt64DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceInt64DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VInt64 = append(a.VInt64, 5)
    a.VInt64 = append(a.VInt64, 3) // sa = {VB, VA}, sb = {VB}
    testSliceInt64DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VInt64 = append(b.VInt64, 5) // sa = {VB, VA}, sb = {VB, VB}
    testSliceInt64DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VInt64 = a.VInt64[:len(a.VInt64)-1] // sa = {VB}, sb = {VB, VB}
    testSliceInt64DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VInt64 = append(a.VInt64, 3)
    a.VInt64 = append(a.VInt64, 3)
    a.VInt64 = append(a.VInt64, 3)
    a.VInt64 = append(a.VInt64, 3) 
    b = SliceTestObject{}
    b.VInt64 = append(b.VInt64, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA}
    testSliceInt64DiffAndApply(t, a, b, 4)
}

func testSliceInt64DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
