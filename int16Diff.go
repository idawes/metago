// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int16Chg struct {
    BaseAttrChg
    OldValue int16
    NewValue int16
}

func NewInt16Chg(s *Attrdef, values ...int16) AttrChg {
    d := Int16Chg{BaseAttrChg: BaseAttrChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
