package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestMapByteInt32(t *testing.T) {
	a := MapTestObject{}
	b := MapTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VByteInt32 = make(map[byte]int32)
    a.VByteInt32[2] =  3 // sa = {K:VA}, sb = nil
    testMapByteInt32DiffAndApply(t, a, b, 1)

 /*   
    // single modification diff at index 0
    b.VByteInt32 = append(b.VByteInt32, 5) // sa = {VA}, sb = {VB}
    testMapByteInt32DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = MapTestObject{} // sa = nil, sb = {VB}
    testMapByteInt32DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VByteInt32 = append(a.VByteInt32, 5)
    a.VByteInt32 = append(a.VByteInt32, 3) // sa = {VB, VA}, sb = {VB}
    testMapByteInt32DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VByteInt32 = append(b.VByteInt32, 5) // sa = {VB, VA}, sb = {VB, VB}
    testMapByteInt32DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VByteInt32 = a.VByteInt32[:len(a.VByteInt32)-1] // sa = {VB}, sb = {VB, VB}
    testMapByteInt32DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VByteInt32 = append(a.VByteInt32, 3)
    a.VByteInt32 = append(a.VByteInt32, 3)
    a.VByteInt32 = append(a.VByteInt32, 3)
    a.VByteInt32 = append(a.VByteInt32, 3) 
    b = MapTestObject{}
    b.VByteInt32 = append(b.VByteInt32, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testMapByteInt32DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VByteInt32[0] = 3
    b.VByteInt32 = append(b.VByteInt32, 5) 
    b.VByteInt32 = append(b.VByteInt32, 5) 
    b.VByteInt32 = append(b.VByteInt32, 5) 
    b.VByteInt32 = append(b.VByteInt32, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteInt32DiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VByteInt32[0] = 3
    a.VByteInt32 = a.VByteInt32[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteInt32DiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VByteInt32[0] = 5
    a.VByteInt32 = append(a.VByteInt32, 3) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
    testMapByteInt32DiffAndApply(t, a, b, 5)
     */
}

func testMapByteInt32DiffAndApply(t *testing.T, a, b MapTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
