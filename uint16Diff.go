// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint16Diff struct {
    BaseAttrDiff
    OldValue uint16
    NewValue uint16
}

func NewUint16Diff(schemaref *Attrdef, values ...uint16) AttrDiff {
    d := Uint16Diff{BaseAttrDiff: BaseAttrDiff{schemaref: schemaref}}
    if len(values) > 0 {
        d.OldValue = values[0]
    }
    if len(values) > 1 {
        d.NewValue = values[1]
    }
    return &d
}
