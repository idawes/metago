// Automatically generated from genericBaseChg.tmpl. DO NOT EDIT!!!!

package metago

import (
    "fmt"
    "io"
)

type Int64Chg struct {
    BaseChg
    OldValue int64
    NewValue int64
}

func (c *Int64Chg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintln(w, "Int64Chg --", c.BaseChg.schemaref, "-- Old:", c.OldValue, ", New:", c.NewValue) 
}

func NewInt64Chg(s *Attrdef, values ...int64) Chg {
    d := Int64Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
