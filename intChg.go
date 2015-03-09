// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type IntChg struct {
    BaseChg
    OldValue int
    NewValue int
}

func NewIntChg(s *Attrdef, values ...int) Chg {
    d := IntChg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
