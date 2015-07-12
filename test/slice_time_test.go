package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestSliceTime(t *testing.T) {
	a := SliceTestObject{}
	b := SliceTestObject{}
	assert.Equal(t, a, b)
	assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\an:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

	// single deletion diff at index 0
	a.VTime = append(a.VTime, time.Unix(1436000000, 0)) // sa = {VA}, sb = nil
	testSliceTimeDiffAndApply(t, a, b, 1)

	// single modification diff at index 0
	b.VTime = append(b.VTime, time.Unix(1436100000, 0)) // sa = {VA}, sb = {VB}
	testSliceTimeDiffAndApply(t, a, b, 1)

	// single insertion diff at index 0
	a = SliceTestObject{} // sa = nil, sb = {VB}
	testSliceTimeDiffAndApply(t, a, b, 1)

	// single deletion diff at index > 0
	a.VTime = append(a.VTime, time.Unix(1436100000, 0))
	a.VTime = append(a.VTime, time.Unix(1436000000, 0)) // sa = {VB, VA}, sb = {VB}
	testSliceTimeDiffAndApply(t, a, b, 1)

	// single modification diff at index > 0
	b.VTime = append(b.VTime, time.Unix(1436100000, 0)) // sa = {VB, VA}, sb = {VB, VB}
	testSliceTimeDiffAndApply(t, a, b, 1)

	// single insertion diff at index > 0
	a.VTime = a.VTime[:len(a.VTime)-1] // sa = {VB}, sb = {VB, VB}
	testSliceTimeDiffAndApply(t, a, b, 1)

	// multiple deletion diff
	a.VTime = append(a.VTime, time.Unix(1436000000, 0))
	a.VTime = append(a.VTime, time.Unix(1436000000, 0))
	a.VTime = append(a.VTime, time.Unix(1436000000, 0))
	a.VTime = append(a.VTime, time.Unix(1436000000, 0))
	b = SliceTestObject{}
	b.VTime = append(b.VTime, time.Unix(1436100000, 0)) // sa = {VB, VA, VA, VA, VA}, sb = {VB}
	testSliceTimeDiffAndApply(t, a, b, 4)

	// multiple modification diff
	b.VTime[0] = time.Unix(1436000000, 0)
	b.VTime = append(b.VTime, time.Unix(1436100000, 0))
	b.VTime = append(b.VTime, time.Unix(1436100000, 0))
	b.VTime = append(b.VTime, time.Unix(1436100000, 0))
	b.VTime = append(b.VTime, time.Unix(1436100000, 0)) // sa = {VB, VA, VA, VA, VA}, sb = {VA, VB, VB, VB, VB}
	testSliceTimeDiffAndApply(t, a, b, 5)

	// multiple insertion diff
	a.VTime[0] = time.Unix(1436000000, 0)
	a.VTime = a.VTime[:1] // sa = {VA}, sb = {VA, VB, VB, VB, VB}
	testSliceTimeDiffAndApply(t, a, b, 4)

	// multiple modifications and insertions diff
	a.VTime[0] = time.Unix(1436100000, 0)
	a.VTime = append(a.VTime, time.Unix(1436000000, 0)) // sa = {VA, VB}, sb = {VA, VB, VB, VB, VB}
	testSliceTimeDiffAndApply(t, a, b, 5)

}

func testSliceTimeDiffAndApply(t *testing.T, a, b SliceTestObject, numChanges int) {
	assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
	assert.NotEqual(t, a, b)

	d := a.Diff(b)
	assert.Equal(t, numChanges, len(d.Chgs), spew.Sdump(d))

	a.Apply(d)
	assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\na:\n%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
	assert.Equal(t, a, b)
}
