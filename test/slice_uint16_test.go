package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestSliceUint16(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VUint16 = append(a.VUint16, 3) // sa = {VA}, sb = nil
    testSliceUint16DiffAndApply(t, a, b, 1)
    
    // single modification diff at index 0
    b.VUint16 = append(b.VUint16, 5) // sa = {VA}, sb = {VB}
    testSliceUint16DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = SliceTestObject{} // sa = nil, sb = {VB}
    testSliceUint16DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VUint16 = append(a.VUint16, 5)
    a.VUint16 = append(a.VUint16, 3) // sa = {VB, VA}, sb = {VB}
    testSliceUint16DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VUint16 = append(b.VUint16, 5) // sa = {VB, VA}, sb = {VB, VB}
    testSliceUint16DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VUint16 = a.VUint16[:len(a.VUint16)-1] // sa = {VB}, sb = {VB, VB}
    testSliceUint16DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VUint16 = append(a.VUint16, 3)
    a.VUint16 = append(a.VUint16, 3)
    a.VUint16 = append(a.VUint16, 3)
    a.VUint16 = append(a.VUint16, 3) 
    b = SliceTestObject{}
    b.VUint16 = append(b.VUint16, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testSliceUint16DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VUint16[0] = 3
    b.VUint16 = append(b.VUint16, 5) 
    b.VUint16 = append(b.VUint16, 5) 
    b.VUint16 = append(b.VUint16, 5) 
    b.VUint16 = append(b.VUint16, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testSliceUint16DiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VUint16[0] = 3
    a.VUint16 = a.VUint16[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testSliceUint16DiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VUint16[0] = 5
    a.VUint16 = append(a.VUint16, 3)
    testSliceUint16DiffAndApply(t, a, b, 5)

}

func testSliceUint16DiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
