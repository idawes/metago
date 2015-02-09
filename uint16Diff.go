// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint16Diff struct {
	BaseAttrChg
	OldValue uint16
	NewValue uint16
}

func NewUint16Diff(schemaref *Attrdef, values ...uint16) AttrChg {
	d := Uint16Diff{BaseAttrChg: BaseAttrChg{schemaref: schemaref}}
	if len(values) > 0 {
		d.OldValue = values[0]
	}
	if len(values) > 1 {
		d.NewValue = values[1]
	}
	return &d
}
