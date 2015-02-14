// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type StringChg struct {
    BaseAttrChg
    OldValue string
    NewValue string
}

func NewStringChg(s *Attrdef, values ...string) AttrChg {
    d := StringChg{BaseAttrChg: BaseAttrChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
