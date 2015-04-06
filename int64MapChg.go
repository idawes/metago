package metago

import (
    "fmt"
    "io"
)


type Int64MapChg struct {
    BaseChg
    Key int64
    Typ ChangeType
    Chgs []Chg
}

func (c *Int64MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Int64MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewInt64MapChg(s *Attrdef, key int64, typ ChangeType, chgs []Chg) Chg {
    return &Int64MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
