// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint32Diff struct {
    BaseAttrDiff
    OldValue uint32
    NewValue uint32
}

func NewUint32Diff(schemaref *Attrdef, values ...uint32) AttrDiff {
    d := Uint32Diff{BaseAttrDiff: BaseAttrDiff{schemaref: schemaref}}
    if len(values) > 0 {
        d.OldValue = values[0]
    }
    if len(values) > 1 {
        d.NewValue = values[1]
    }
    return &d
}
