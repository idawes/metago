package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestBasicTime(t *testing.T) {
	a := BasicAttrTypesObject{}
	b := BasicAttrTypesObject{}

	a.VTime = time.Now()
	b.VTime = time.Now().Add(5 * time.Hour)
	assert.NotEqual(t, a, b)
	assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

	d := a.Diff(b)
	assert.Equal(t, len(d.Chgs), 1, spew.Sdump(d))

	a.Apply(d)
	assert.Equal(t, a, b)
	assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\nn:\a%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))
}
