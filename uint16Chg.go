// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint16Chg struct {
    BaseChg
    OldValue uint16
    NewValue uint16
}

func NewUint16Chg(s *Attrdef, values ...uint16) Chg {
    d := Uint16Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
