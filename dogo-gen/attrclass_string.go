// generated by stringer -type=attrClass; DO NOT EDIT

package main

import "fmt"

const _attrClass_name = "ATTR_CLASS_BUILTINATTR_CLASS_QUALIFIED_BUILTINATTR_CLASS_SLICEATTR_CLASS_MAPATTR_CLASS_DIFFABLE_OBJ"

var _attrClass_index = [...]uint8{18, 46, 62, 76, 99}

func (i attrClass) String() string {
	if i < 0 || i >= attrClass(len(_attrClass_index)) {
		return fmt.Sprintf("attrClass(%d)", i)
	}
	hi := _attrClass_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _attrClass_index[i-1]
	}
	return _attrClass_name[lo:hi]
}
