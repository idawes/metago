package metago

import (
    "fmt"
    "io"
)


type Int32MapChg struct {
    BaseChg
    Key int32
    Typ ChangeType
    Chgs []Chg
}

func (c *Int32MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Int32MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewInt32MapChg(s *Attrdef, key int32, typ ChangeType, chgs []Chg) Chg {
    return &Int32MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
