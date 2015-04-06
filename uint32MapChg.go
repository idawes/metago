package metago

import (
    "fmt"
    "io"
)


type Uint32MapChg struct {
    BaseChg
    Key uint32
    Typ ChangeType
    Chgs []Chg
}

func (c *Uint32MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Uint32MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewUint32MapChg(s *Attrdef, key uint32, typ ChangeType, chgs []Chg) Chg {
    return &Uint32MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
