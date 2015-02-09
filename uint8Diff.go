// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint8Diff struct {
	BaseAttrChg
	OldValue uint8
	NewValue uint8
}

func NewUint8Diff(schemaref *Attrdef, values ...uint8) AttrChg {
	d := Uint8Diff{BaseAttrChg: BaseAttrChg{schemaref: schemaref}}
	if len(values) > 0 {
		d.OldValue = values[0]
	}
	if len(values) > 1 {
		d.NewValue = values[1]
	}
	return &d
}
