// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint8Diff struct {
    BaseAttrDiff
    OldValue uint8
    NewValue uint8
}

func NewUint8Diff(schemaref *Attrdef, values ...uint8) AttrDiff {
    d := Uint8Diff{BaseAttrDiff: BaseAttrDiff{schemaref: schemaref}}
    if len(values) > 0 {
        d.OldValue = values[0]
    }
    if len(values) > 1 {
        d.NewValue = values[1]
    }
    return &d
}
