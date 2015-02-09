// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int64Diff struct {
    BaseAttrDiff
    OldValue int64
    NewValue int64
}

func NewInt64Diff(schemaref *Attrdef, values ...int64) AttrDiff {
    d := Int64Diff{BaseAttrDiff: BaseAttrDiff{schemaref: schemaref}}
    if len(values) > 0 {
        d.OldValue = values[0]
    }
    if len(values) > 1 {
        d.NewValue = values[1]
    }
    return &d
}
