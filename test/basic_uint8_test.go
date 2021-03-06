package test

import (
    "fmt"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/davecgh/go-spew/spew"
)

func TestBasicUint8(t *testing.T) {
	a := BasicAttrTypesObject{}
	b := BasicAttrTypesObject{}
    
    a.VUint8 = 3 
    b.VUint8 = 5
    assert.NotEqual(t, a, b)
    assert.Equal(t, a.Equals(b), false, fmt.Sprintf("\na:\n%s\nb:\n%s\n", spew.Sdump(a), spew.Sdump(b)))

	d := a.Diff(b)
    assert.Equal(t, len(d.Chgs), 1, spew.Sdump(d))

	a.Apply(d)
    assert.Equal(t, a, b)
    assert.Equal(t, a.Equals(b), true, fmt.Sprintf("\nn:\a%s\nb:\n%s\ndiff:\n%s", spew.Sdump(a), spew.Sdump(b), spew.Sdump(d)))
}
