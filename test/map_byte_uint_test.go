package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestMapByteUint(t *testing.T) {
	a := MapTestObject{}
	b := MapTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VByteUint = make(map[byte]uint)
    a.VByteUint[2] =  3 // sa = {K:VA}, sb = nil
    testMapByteUintDiffAndApply(t, a, b, 1)

 /*   
    // single modification diff at index 0
    b.VByteUint = append(b.VByteUint, 5) // sa = {VA}, sb = {VB}
    testMapByteUintDiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = MapTestObject{} // sa = nil, sb = {VB}
    testMapByteUintDiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VByteUint = append(a.VByteUint, 5)
    a.VByteUint = append(a.VByteUint, 3) // sa = {VB, VA}, sb = {VB}
    testMapByteUintDiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VByteUint = append(b.VByteUint, 5) // sa = {VB, VA}, sb = {VB, VB}
    testMapByteUintDiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VByteUint = a.VByteUint[:len(a.VByteUint)-1] // sa = {VB}, sb = {VB, VB}
    testMapByteUintDiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VByteUint = append(a.VByteUint, 3)
    a.VByteUint = append(a.VByteUint, 3)
    a.VByteUint = append(a.VByteUint, 3)
    a.VByteUint = append(a.VByteUint, 3) 
    b = MapTestObject{}
    b.VByteUint = append(b.VByteUint, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testMapByteUintDiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VByteUint[0] = 3
    b.VByteUint = append(b.VByteUint, 5) 
    b.VByteUint = append(b.VByteUint, 5) 
    b.VByteUint = append(b.VByteUint, 5) 
    b.VByteUint = append(b.VByteUint, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteUintDiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VByteUint[0] = 3
    a.VByteUint = a.VByteUint[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteUintDiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VByteUint[0] = 5
    a.VByteUint = append(a.VByteUint, 3) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
    testMapByteUintDiffAndApply(t, a, b, 5)
     */
}

func testMapByteUintDiffAndApply(t *testing.T, a, b MapTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
