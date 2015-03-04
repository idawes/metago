// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint8Chg struct {
    BaseChg
    OldValue uint8
    NewValue uint8
}

func NewUint8Chg(s *Attrdef, values ...uint8) Chg {
    d := Uint8Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
