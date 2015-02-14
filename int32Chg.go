// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int32Chg struct {
    BaseAttrChg
    OldValue int32
    NewValue int32
}

func NewInt32Chg(s *Attrdef, values ...int32) AttrChg {
    d := Int32Chg{BaseAttrChg: BaseAttrChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
