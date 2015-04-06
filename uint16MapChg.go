package metago

import (
    "fmt"
    "io"
)


type Uint16MapChg struct {
    BaseChg
    Key uint16
    Typ ChangeType
    Chgs []Chg
}

func (c *Uint16MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Uint16MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewUint16MapChg(s *Attrdef, key uint16, typ ChangeType, chgs []Chg) Chg {
    return &Uint16MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
