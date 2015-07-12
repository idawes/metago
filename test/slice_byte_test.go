package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceByte(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VByte = append(a.VByte, 3) // sa = {VA}, sb = nil
    testSliceByteDiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VByte = append(b.VByte, 5) // sa = {VA}, sb = {VB}
    testSliceByteDiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceByteDiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VByte = append(a.VByte, 5)
    a.VByte = append(a.VByte, 3) // sa = {VB, VA}, sb = {VB}
    testSliceByteDiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VByte = append(b.VByte, 5) // sa = {VB, VA}, sb = {VB, VB}
    testSliceByteDiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VByte = a.VByte[:len(a.VByte)-1] // sa = {VB}, sb = {VB, VB}
    testSliceByteDiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VByte = append(a.VByte, 3)
    a.VByte = append(a.VByte, 3)
    a.VByte = append(a.VByte, 3)
    a.VByte = append(a.VByte, 3) 
    b = SliceTestObject{}
    b.VByte = append(b.VByte, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testSliceByteDiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VByte[0] = 3
    b.VByte = append(b.VByte, 5) 
    b.VByte = append(b.VByte, 5) 
    b.VByte = append(b.VByte, 5) 
    b.VByte = append(b.VByte, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testSliceByteDiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VByte[0] = 3
    a.VByte = a.VByte[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testSliceByteDiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VByte[0] = 5
    a.VByte = append(a.VByte, 3) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
    testSliceByteDiffAndApply(t, a, b, 5)

}

func testSliceByteDiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
