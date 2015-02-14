// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int64Chg struct {
    BaseAttrChg
    OldValue int64
    NewValue int64
}

func NewInt64Chg(s *Attrdef, values ...int64) AttrChg {
    d := Int64Chg{BaseAttrChg: BaseAttrChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
