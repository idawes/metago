package metago

import (
	"bytes"

	"github.com/nu7hatch/gouuid"
)

// TypeId is a globally unique identifier for a metago object type
type TypeID struct {
	pkg *uuid.UUID
	typ int
}

func (t *TypeID) Equals(o *TypeID) bool {
	if t.pkg != o.pkg {
		return false
	}
	if t.typ != o.typ {
		return false
	}
	return true
}

func (t *TypeID) Compare(o *TypeID) int {
	v := bytes.Compare(t.pkg[:], o.pkg[:])
	if v != 0 {
		return v
	}
	return t.typ - o.typ
}

// AttrID is a globally unique identifier for an attribute in a metago object type
type AttrID struct {
	TypeID
	attr int
}

func (a *AttrID) Equals(o *AttrID) bool {
	if a.attr != o.attr {
		return false
	}
	if !a.TypeID.Equals(&o.TypeID) {
		return false
	}
	return true
}

func (a *AttrID) Compare(o *AttrID) int {
	v := a.TypeID.Compare(&o.TypeID)
	if v != 0 {
		return v
	}
	return a.attr - o.attr
}
