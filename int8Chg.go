// Automatically generated from genericBaseChg.tmpl. DO NOT EDIT!!!!

package metago

import (
    "fmt"
    "io"
)

type Int8Chg struct {
    BaseChg
    OldValue int8
    NewValue int8
}

func (c *Int8Chg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintln(w, "Int8Chg --", c.BaseChg.schemaref, "-- Old:", c.OldValue, ", New:", c.NewValue) 
}

func NewInt8Chg(s *Attrdef, values ...int8) Chg {
    d := Int8Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
