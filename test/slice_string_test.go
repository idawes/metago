package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceString(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VString = append(a.VString, "Foo") // sa = {VA}, sb = nil
    testSliceStringDiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VString = append(b.VString, "Bar") // sa = {VA}, sb = {VB}
    testSliceStringDiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceStringDiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VString = append(a.VString, "Bar")
    a.VString = append(a.VString, "Foo") // sa = {VB, VA}, sb = {VB}
    testSliceStringDiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VString = append(b.VString, "Bar") // sa = {VB, VA}, sb = {VB, VB}
    testSliceStringDiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VString = a.VString[:len(a.VString)-1] // sa = {VB}, sb = {VB, VB}
    testSliceStringDiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VString = append(a.VString, "Foo")
    a.VString = append(a.VString, "Foo")
    a.VString = append(a.VString, "Foo")
    a.VString = append(a.VString, "Foo") 
    b = SliceTestObject{}
    b.VString = append(b.VString, "Bar") // sa = {VB, VA, VA, VA, VA}, sb = {VA}
    testSliceStringDiffAndApply(t, a, b, 4)
}

func testSliceStringDiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
