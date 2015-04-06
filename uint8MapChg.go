package metago

import (
    "fmt"
    "io"
)


type Uint8MapChg struct {
    BaseChg
    Key uint8
    Typ ChangeType
    Chgs []Chg
}

func (c *Uint8MapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "Uint8MapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewUint8MapChg(s *Attrdef, key uint8, typ ChangeType, chgs []Chg) Chg {
    return &Uint8MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
