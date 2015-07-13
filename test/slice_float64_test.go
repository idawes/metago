package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceFloat64(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VFloat64 = append(a.VFloat64, 3.23) // sa = {VA}, sb = nil
    testSliceFloat64DiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VFloat64 = append(b.VFloat64, 5.332) // sa = {VA}, sb = {VB}
    testSliceFloat64DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceFloat64DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VFloat64 = append(a.VFloat64, 5.332)
    a.VFloat64 = append(a.VFloat64, 3.23) // sa = {VB, VA}, sb = {VB}
    testSliceFloat64DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VFloat64 = append(b.VFloat64, 5.332) // sa = {VB, VA}, sb = {VB, VB}
    testSliceFloat64DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VFloat64 = a.VFloat64[:len(a.VFloat64)-1] // sa = {VB}, sb = {VB, VB}
    testSliceFloat64DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VFloat64 = append(a.VFloat64, 3.23)
    a.VFloat64 = append(a.VFloat64, 3.23)
    a.VFloat64 = append(a.VFloat64, 3.23)
    a.VFloat64 = append(a.VFloat64, 3.23) 
    b = SliceTestObject{}
    b.VFloat64 = append(b.VFloat64, 5.332) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testSliceFloat64DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VFloat64[0] = 3.23
    b.VFloat64 = append(b.VFloat64, 5.332) 
    b.VFloat64 = append(b.VFloat64, 5.332) 
    b.VFloat64 = append(b.VFloat64, 5.332) 
    b.VFloat64 = append(b.VFloat64, 5.332) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testSliceFloat64DiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VFloat64[0] = 3.23
    a.VFloat64 = a.VFloat64[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testSliceFloat64DiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VFloat64[0] = 5.332
    a.VFloat64 = append(a.VFloat64, 3.23) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
    testSliceFloat64DiffAndApply(t, a, b, 5)
}

func testSliceFloat64DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
