// Automatically generated from genericBaseChg.tmpl. DO NOT EDIT!!!!

package metago

import (
    "fmt"
    "io"
)

type Float64Chg struct {
    BaseChg
    OldValue float64
    NewValue float64
}

func (c *Float64Chg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
    fmt.Fprintln(w, "Float64Chg --", c.BaseChg.schemaref, "-- Old:", c.OldValue, ", New:", c.NewValue) 
}

func NewFloat64Chg(s *Attrdef, values ...float64) Chg {
    d := Float64Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
