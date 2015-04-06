package metago

import (
    "fmt"
    "io"
)


type Int8MapChg struct {
    BaseChg
    Key int8
    Typ ChangeType
    Chgs []Chg
}

func (c *Int8MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Int8MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewInt8MapChg(s *Attrdef, key int8, typ ChangeType, chgs []Chg) Chg {
    return &Int8MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
