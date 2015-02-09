// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int8Diff struct {
	BaseAttrChg
	OldValue int8
	NewValue int8
}

func NewInt8Diff(schemaref *Attrdef, values ...int8) AttrChg {
	d := Int8Diff{BaseAttrChg: BaseAttrChg{schemaref: schemaref}}
	if len(values) > 0 {
		d.OldValue = values[0]
	}
	if len(values) > 1 {
		d.NewValue = values[1]
	}
	return &d
}
