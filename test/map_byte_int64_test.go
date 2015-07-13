package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestMapByteInt64(t *testing.T) {
	a := MapTestObject{}
	b := MapTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VByteInt64 = make(map[byte]int64)
    a.VByteInt64[2] =  3 // sa = {K:VA}, sb = nil
    testMapByteInt64DiffAndApply(t, a, b, 1)

 /*   
    // single modification diff at index 0
    b.VByteInt64 = append(b.VByteInt64, 5) // sa = {VA}, sb = {VB}
    testMapByteInt64DiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = MapTestObject{} // sa = nil, sb = {VB}
    testMapByteInt64DiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VByteInt64 = append(a.VByteInt64, 5)
    a.VByteInt64 = append(a.VByteInt64, 3) // sa = {VB, VA}, sb = {VB}
    testMapByteInt64DiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VByteInt64 = append(b.VByteInt64, 5) // sa = {VB, VA}, sb = {VB, VB}
    testMapByteInt64DiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VByteInt64 = a.VByteInt64[:len(a.VByteInt64)-1] // sa = {VB}, sb = {VB, VB}
    testMapByteInt64DiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VByteInt64 = append(a.VByteInt64, 3)
    a.VByteInt64 = append(a.VByteInt64, 3)
    a.VByteInt64 = append(a.VByteInt64, 3)
    a.VByteInt64 = append(a.VByteInt64, 3) 
    b = MapTestObject{}
    b.VByteInt64 = append(b.VByteInt64, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testMapByteInt64DiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VByteInt64[0] = 3
    b.VByteInt64 = append(b.VByteInt64, 5) 
    b.VByteInt64 = append(b.VByteInt64, 5) 
    b.VByteInt64 = append(b.VByteInt64, 5) 
    b.VByteInt64 = append(b.VByteInt64, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteInt64DiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VByteInt64[0] = 3
    a.VByteInt64 = a.VByteInt64[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteInt64DiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VByteInt64[0] = 5
    a.VByteInt64 = append(a.VByteInt64, 3) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
    testMapByteInt64DiffAndApply(t, a, b, 5)
     */
}

func testMapByteInt64DiffAndApply(t *testing.T, a, b MapTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
