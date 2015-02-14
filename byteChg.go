// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type ByteChg struct {
    BaseAttrChg
    OldValue byte
    NewValue byte
}

func NewByteChg(s *Attrdef, values ...byte) AttrChg {
    d := ByteChg{BaseAttrChg: BaseAttrChg{schemaref: s}}
    if len(values) > 0 {
        d.NewValue = values[0]
    }
    if len(values) > 1 {
        d.OldValue = values[1]
    }
    return &d
}
