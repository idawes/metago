package metago

import (
	"fmt"
	"io"
	"time"
)

type TimeMapChg struct {
	BaseChg
	Key  time.Time
	Typ  ChangeType
	Chgs []Chg
}

func (c *TimeMapChg) WriteIndented(w io.Writer, lev int) {
	for i := 0; i < lev; i++ {
		fmt.Fprintf(w, "  ")
	}
	fmt.Fprintf(w, "TimeMapChg (%s) -- %s -- Key: %s", c.Typ, c.BaseChg.schemaref, c.Key)
	lev++
	for _, c1 := range c.Chgs {
		c1.WriteIndented(w, lev)
	}
}

func NewTimeMapChg(s *Attrdef, key time.Time, typ ChangeType, chgs []Chg) Chg {
	return &TimeMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
