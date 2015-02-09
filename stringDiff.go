// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type StringDiff struct {
	BaseAttrChg
	OldValue string
	NewValue string
}

func NewStringDiff(schemaref *Attrdef, values ...string) AttrChg {
	d := StringDiff{BaseAttrChg: BaseAttrChg{schemaref: schemaref}}
	if len(values) > 0 {
		d.OldValue = values[0]
	}
	if len(values) > 1 {
		d.NewValue = values[1]
	}
	return &d
}
