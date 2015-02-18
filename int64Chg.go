// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int64Chg struct {
	BaseChg
	OldValue int64
	NewValue int64
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
