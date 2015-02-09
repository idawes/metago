// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Uint64Diff struct {
    BaseAttrDiff
    OldValue uint64
    NewValue uint64
}

func NewUint64Diff(schemaref *Attrdef, values ...uint64) AttrDiff {
    d := Uint64Diff{BaseAttrDiff: BaseAttrDiff{schemaref: schemaref}}
    if len(values) > 0 {
        d.OldValue = values[0]
    }
    if len(values) > 1 {
        d.NewValue = values[1]
    }
    return &d
}
