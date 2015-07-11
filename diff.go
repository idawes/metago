package metago

import (
	"bytes"
	"fmt"
	"io"
)

type Diff struct {
	Chgs []Chg
}

func (d Diff) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
	fmt.Fprintf(w, "Diff: \n")
	lev++
	for _, c := range d.Chgs {
		c.WriteIndented(w, lev)
	}
}

type Chg interface {
	AttributeID() *AttrID
	Schemaref() *Attrdef
	WriteIndented(w io.Writer, lev int)
}

type BaseChg struct {
	schemaref *Attrdef
}

func (d BaseChg) AttributeID() *AttrID {
	return d.schemaref.ID
}

func (d BaseChg) Schemaref() *Attrdef {
	return d.schemaref
}

func (d BaseChg) PersistenceClass() PersistenceClass {
	return d.schemaref.Persistence
}

func (d BaseChg) WriteTo(w *Writer) error {
	w.Write(d.schemaref.ID.Pkg[:])
	w.WriteVarint(int64(d.schemaref.ID.Typ))
	w.WriteVarint(int64(d.schemaref.ID.Attr))
	return nil
}

//go:generate stringer -type=ChangeType
type ChangeType int

const (
	ChangeTypeInsert ChangeType = iota
	ChangeTypeDelete
	ChangeTypeModify
)

type SliceChg struct {
	BaseChg
	Idx  int
	Typ  ChangeType
	Chgs []Chg
}

func (c SliceChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
	fmt.Fprintf(w, "SliceChg (%s) -- %s -- Idx: %d\n", c.Typ, c.BaseChg.schemaref, c.Idx)
	lev++
	for _, c1 := range c.Chgs {
		c1.WriteIndented(w, lev)
	}
}

func NewSliceChg(sref *Attrdef, idx int, typ ChangeType, chgs []Chg) Chg {
	return &SliceChg{BaseChg: BaseChg{schemaref: sref}, Idx: idx, Typ: typ, Chgs: chgs}
}

type StructChg struct {
	BaseChg
	Chg Diff
}

func (c StructChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
	fmt.Fprintf(w, "StructChg: ")
	c.Chg.WriteIndented(w, lev+1)
}

func NewStructChg(sref *Attrdef, chg Diff) Chg {
	return &StructChg{BaseChg: BaseChg{schemaref: sref}, Chg: chg}
}

type DiffApplyError struct {
	errChain []error
}

func (e DiffApplyError) Error() string {
	buf := &bytes.Buffer{}
	for i, err := range e.errChain {
		if i != 0 {
			fmt.Fprintln(buf, " Caused by:")
		}
		for j := 0; j < i; j++ {
			fmt.Fprint(buf, "  ")
		}
		fmt.Fprint(buf, err)
	}
	fmt.Fprintln(buf)
	return buf.String()
}
