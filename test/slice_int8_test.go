package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceInt8(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VInt8 = append(a.VInt8, 3) // sa = {VA}, sb = nil
    testSliceInt8DiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VInt8 = append(b.VInt8, 5) // sa = {VA}, sb = {VB}
    testSliceInt8DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceInt8DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VInt8 = append(a.VInt8, 5)
    a.VInt8 = append(a.VInt8, 3) // sa = {VB, VA}, sb = {VB}
    testSliceInt8DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VInt8 = append(b.VInt8, 5) // sa = {VB, VA}, sb = {VB, VB}
    testSliceInt8DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VInt8 = a.VInt8[:len(a.VInt8)-1] // sa = {VB}, sb = {VB, VB}
    testSliceInt8DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VInt8 = append(a.VInt8, 3)
    a.VInt8 = append(a.VInt8, 3)
    a.VInt8 = append(a.VInt8, 3)
    a.VInt8 = append(a.VInt8, 3) 
    b = SliceTestObject{}
    b.VInt8 = append(b.VInt8, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testSliceInt8DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VInt8[0] = 3
    b.VInt8 = append(b.VInt8, 5) 
    b.VInt8 = append(b.VInt8, 5) 
    b.VInt8 = append(b.VInt8, 5) 
    b.VInt8 = append(b.VInt8, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testSliceInt8DiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VInt8[0] = 3
    a.VInt8 = a.VInt8[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testSliceInt8DiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VInt8[0] = 5
    a.VInt8 = append(a.VInt8, 3) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
    testSliceInt8DiffAndApply(t, a, b, 5)
}

func testSliceInt8DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
