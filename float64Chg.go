// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Float64Chg struct {
    BaseChg
    OldValue float64
    NewValue float64
}

func NewFloat64Chg(s *Attrdef, values ...float64) Chg {
    d := Float64Chg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
