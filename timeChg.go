// Automatically generated from genericBaseChg.tmpl. DO NOT EDIT!!!!

package metago

import (
	"fmt"
	"io"
	"time"
)

type TimeChg struct {
	BaseChg
	OldValue time.Time
	NewValue time.Time
}

func (c *TimeChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
	fmt.Fprintln(w, "TimeChg --", c.BaseChg.schemaref, "-- Old:", c.OldValue, ", New:", c.NewValue)
}

func NewTimeChg(s *Attrdef, values ...time.Time) Chg {
	d := TimeChg{BaseChg: BaseChg{schemaref: s}}
	if len(values) > 0 {
		d.NewValue = values[0]
	}
	if len(values) > 1 {
		d.OldValue = values[1]
	}
	return &d
}
