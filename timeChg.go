// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

import "time"

type TimeChg struct {
	BaseAttrChg
	OldValue time.Time
	NewValue time.Time
}

func NewTimeChg(s *Attrdef, values ...time.Time) AttrChg {
	d := TimeChg{BaseAttrChg: BaseAttrChg{schemaref: s}}
	if len(values) > 0 {
		d.NewValue = values[0]
	}
	if len(values) > 1 {
		d.OldValue = values[1]
	}
	return &d
}
