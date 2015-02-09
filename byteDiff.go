// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type ByteDiff struct {
	BaseAttrChg
	OldValue byte
	NewValue byte
}

func NewByteDiff(schemaref *Attrdef, values ...byte) AttrChg {
	d := ByteDiff{BaseAttrChg: BaseAttrChg{schemaref: schemaref}}
	if len(values) > 0 {
		d.OldValue = values[0]
	}
	if len(values) > 1 {
		d.NewValue = values[1]
	}
	return &d
}
