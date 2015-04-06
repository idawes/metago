package metago

import (
    "fmt"
    "io"
)


type Int16MapChg struct {
    BaseChg
    Key int16
    Typ ChangeType
    Chgs []Chg
}

func (c *Int16MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Int16MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewInt16MapChg(s *Attrdef, key int16, typ ChangeType, chgs []Chg) Chg {
    return &Int16MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
