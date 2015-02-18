// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int8Chg struct {
	BaseChg
	OldValue int8
	NewValue int8
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
