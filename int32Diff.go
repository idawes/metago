// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int32Diff struct {
	BaseAttrChg
	OldValue int32
	NewValue int32
}

func NewInt32Diff(schemaref *Attrdef, values ...int32) AttrChg {
	d := Int32Diff{BaseAttrChg: BaseAttrChg{schemaref: schemaref}}
	if len(values) > 0 {
		d.OldValue = values[0]
	}
	if len(values) > 1 {
		d.NewValue = values[1]
	}
	return &d
}
