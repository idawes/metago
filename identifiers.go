package metago

import (
	"bytes"

	"github.com/nu7hatch/gouuid"
)

// TypeId is a globally unique identifier for a metago object type
type TypeID struct {
	Pkg *uuid.UUID
	Typ int
}

func (t *TypeID) Equals(o *TypeID) bool {
	if t.Pkg != o.Pkg {
		return false
	}
	if t.Typ != o.Typ {
		return false
	}
	return true
}

func (t *TypeID) Compare(o *TypeID) int {
	v := bytes.Compare(t.Pkg[:], o.Pkg[:])
	if v != 0 {
		return v
	}
	return t.Typ - o.Typ
}

// AttrID is a globally unique identifier for an attribute in a metago object type
type AttrID struct {
	TypeID
	Attr int
}

func (a *AttrID) Equals(o *AttrID) bool {
	if a.Attr != o.Attr {
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
	return a.Attr - o.Attr
}
