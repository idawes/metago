// Automatically generated from generic_diff.tmpl. DO NOT EDIT!!!!

package metago

type Int16Diff struct {
    BaseAttrDiff
    OldValue int16
    NewValue int16
}

func NewInt16Diff(schemaref *Attrdef, values ...int16) AttrDiff {
    d := Int16Diff{BaseAttrDiff: BaseAttrDiff{schemaref: schemaref}}
    if len(values) > 0 {
        d.OldValue = values[0]
    }
    if len(values) > 1 {
        d.NewValue = values[1]
    }
    return &d
}
