// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type UintChg struct {
    BaseChg
    OldValue uint
    NewValue uint
}

func NewUintChg(s *Attrdef, values ...uint) Chg {
    d := UintChg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
