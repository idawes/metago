// Automatically generated from genericBaseChg.tmpl. DO NOT EDIT!!!!

package metago

import (
    "fmt"
    "io"
)

type Uint8Chg struct {
    BaseChg
    OldValue uint8
    NewValue uint8
}

func (c *Uint8Chg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintln(w, "Uint8Chg --", c.BaseChg.schemaref, "-- Old:", c.OldValue, ", New:", c.NewValue) 
}

func NewUint8Chg(s *Attrdef, values ...uint8) Chg {
    d := Uint8Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
