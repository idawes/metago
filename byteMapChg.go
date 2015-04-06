package metago

import (
    "fmt"
    "io"
)


type ByteMapChg struct {
    BaseChg
    Key byte
    Typ ChangeType
    Chgs []Chg
}

func (c *ByteMapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintf(w, "ByteMapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key) 
    lev++
    for _, c1 := range c.Chgs {
       c1.WriteIndented(w, lev)
    }
}

func NewByteMapChg(s *Attrdef, key byte, typ ChangeType, chgs []Chg) Chg {
    return &ByteMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
