// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint32Chg struct {
	BaseChg
	OldValue uint32
	NewValue uint32
}

func NewUint32Chg(s *Attrdef, values ...uint32) Chg {
	d := Uint32Chg{BaseChg: BaseChg{schemaref: s}}
	if len(values) > 0 {
		d.NewValue = values[0]
	}
	if len(values) > 1 {
		d.OldValue = values[1]
	}
	return &d
}
