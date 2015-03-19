// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Float32Chg struct {
    BaseChg
    OldValue float32
    NewValue float32
}

func NewFloat32Chg(s *Attrdef, values ...float32) Chg {
    d := Float32Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
