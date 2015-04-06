// Automatically generated from genericBaseChg.tmpl. DO NOT EDIT!!!!

package metago

import (
    "fmt"
    "io"
)

type ByteChg struct {
    BaseChg
    OldValue byte
    NewValue byte
}

func (c *ByteChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintln(w, "ByteChg --", c.BaseChg.schemaref, "-- Old:", c.OldValue, ", New:", c.NewValue) 
}

func NewByteChg(s *Attrdef, values ...byte) Chg {
    d := ByteChg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
