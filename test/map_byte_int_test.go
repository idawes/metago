package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestMapByteInt(t *testing.T) {
	a := MapTestObject{}
	b := MapTestObject{}
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

    // single deletion diff at index 0
    a.VByteInt = make(map[byte]int)
    a.VByteInt[2] =  3 // sa = {K:VA}, sb = nil
    testMapByteIntDiffAndApply(t, a, b, 1)

 /*   
    // single modification diff at index 0
    b.VByteInt = append(b.VByteInt, 5) // sa = {VA}, sb = {VB}
    testMapByteIntDiffAndApply(t, a, b, 1)

    // single insertion diff at index 0
    a = MapTestObject{} // sa = nil, sb = {VB}
    testMapByteIntDiffAndApply(t, a, b, 1)

    // single deletion diff at index > 0 
    a.VByteInt = append(a.VByteInt, 5)
    a.VByteInt = append(a.VByteInt, 3) // sa = {VB, VA}, sb = {VB}
    testMapByteIntDiffAndApply(t, a, b, 1)

    // single modification diff at index > 0
    b.VByteInt = append(b.VByteInt, 5) // sa = {VB, VA}, sb = {VB, VB}
    testMapByteIntDiffAndApply(t, a, b, 1)
    
    // single insertion diff at index > 0
    a.VByteInt = a.VByteInt[:len(a.VByteInt)-1] // sa = {VB}, sb = {VB, VB}
    testMapByteIntDiffAndApply(t, a, b, 1)

    // multiple deletion diff
    a.VByteInt = append(a.VByteInt, 3)
    a.VByteInt = append(a.VByteInt, 3)
    a.VByteInt = append(a.VByteInt, 3)
    a.VByteInt = append(a.VByteInt, 3) 
    b = MapTestObject{}
    b.VByteInt = append(b.VByteInt, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
    testMapByteIntDiffAndApply(t, a, b, 4)

    // multiple modification diff
    b.VByteInt[0] = 3
    b.VByteInt = append(b.VByteInt, 5) 
    b.VByteInt = append(b.VByteInt, 5) 
    b.VByteInt = append(b.VByteInt, 5) 
    b.VByteInt = append(b.VByteInt, 5) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteIntDiffAndApply(t, a, b, 5)
    
    // multiple insertion diff
    a.VByteInt[0] = 3
    a.VByteInt = a.VByteInt[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
    testMapByteIntDiffAndApply(t, a, b, 4)

    // multiple modifications and insertions diff
    a.VByteInt[0] = 5
    a.VByteInt = append(a.VByteInt, 3) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
    testMapByteIntDiffAndApply(t, a, b, 5)
     */
}

func testMapByteIntDiffAndApply(t *testing.T, a, b MapTestObject, numChanges int) {
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
    assert.NotEqual(t, a, b)

	d := a.Diff(b)
    assert.Equal(t, numChanges , len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
    assert.Equal(t, a, b)
}
