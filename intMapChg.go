package metago

import (
    "fmt"
    "io"
)


type IntMapChg struct {
    BaseChg
    Key int
    Typ ChangeType
    Chgs []Chg
}

func (c *IntMapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "IntMapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewIntMapChg(s *Attrdef, key int, typ ChangeType, chgs []Chg) Chg {
    return &IntMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
