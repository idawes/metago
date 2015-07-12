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

	// make a longer than b
	a.VTime = append(a.VTime, time.Unix(1436000000, 0))
	testSliceTimeDiffAndApply(t, a, b, 1)

	// make a and b the same length, but with a change.
	b.VTime = append(b.VTime, time.Unix(1436100000, 0))
	testSliceTimeDiffAndApply(t, a, b, 1)

	// make a shorter than b
	a = SliceTestObject{}
	testSliceTimeDiffAndApply(t, a, b, 1)

	// make both non-nil, and a longer than b
	a.VTime = append(a.VTime, time.Unix(1436100000, 0))
	a.VTime = append(a.VTime, time.Unix(1436000000, 0))
	testSliceTimeDiffAndApply(t, a, b, 1)

	// make both same length, but with a change
	b.VTime = append(b.VTime, time.Unix(1436100000, 0))
	testSliceTimeDiffAndApply(t, a, b, 1)

	// make both non-nil, and a shorter than b
	a.VTime = a.VTime[:len(a.VTime)-1]
	testSliceTimeDiffAndApply(t, a, b, 1)

	// make 2 changes
	a.VTime[0] = time.Unix(1436000000, 0)
	testSliceTimeDiffAndApply(t, a, b, 2)
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
