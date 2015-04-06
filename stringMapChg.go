package metago

import (
    "fmt"
    "io"
)


type StringMapChg struct {
    BaseChg
    Key string
    Typ ChangeType
    Chgs []Chg
}

func (c *StringMapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "StringMapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewStringMapChg(s *Attrdef, key string, typ ChangeType, chgs []Chg) Chg {
    return &StringMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
