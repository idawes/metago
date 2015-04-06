package metago

import (
    "fmt"
    "io"
)


type Uint64MapChg struct {
    BaseChg
    Key uint64
    Typ ChangeType
    Chgs []Chg
}

func (c *Uint64MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Uint64MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewUint64MapChg(s *Attrdef, key uint64, typ ChangeType, chgs []Chg) Chg {
    return &Uint64MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
