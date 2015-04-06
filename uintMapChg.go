package metago

import (
    "fmt"
    "io"
)


type UintMapChg struct {
    BaseChg
    Key uint
    Typ ChangeType
    Chgs []Chg
}

func (c *UintMapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "UintMapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewUintMapChg(s *Attrdef, key uint, typ ChangeType, chgs []Chg) Chg {
    return &UintMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
