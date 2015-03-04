// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type ByteChg struct {
    BaseChg
    OldValue byte
    NewValue byte
}

func NewByteChg(s *Attrdef, values ...byte) Chg {
    d := ByteChg{BaseChg: BaseChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
