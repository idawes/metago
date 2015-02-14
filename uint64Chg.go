// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint64Chg struct {
    BaseAttrChg
    OldValue uint64
    NewValue uint64
}

func NewUint64Chg(s *Attrdef, values ...uint64) AttrChg {
    d := Uint64Chg{BaseAttrChg: BaseAttrChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
