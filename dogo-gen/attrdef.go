package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	attribute definition:
	<id> <name> <type> [<persistenceType>]
*/

func parseAttribute(t *typedef, r *reader) (attrDef, error) {
	fields, err := r.read()
	if len(fields) < 3 || len(fields) > 4 {
		return nil, fmt.Errorf("invalid attribute specification, line %d of file %s", r.line, r.f.Name())
	}
	id, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("expecting an integer attribute id, found \"%s\", line %d of file %s", fields[0], r.line, r.f.Name())
	}
	name := fields[1]
	p := persistenceClassPersistent
	if len(fields) > 3 {
		switch fields[3] {
		case "persistent":
			p = persistenceClassPersistent
		case "non-persistent":
			p = persistenceClassNonPersistent
		case "ephemeral":
			p = persistenceClassEphemeral
		default:
			return nil, fmt.Errorf("unrecognized persistence type \"%s\", line %d of file %s", fields[3], r.line, r.f.Name())
		}
	}
	return newAttrDef(&baseAttrDef{parentType: t, attributeID: id, name: name, attrType: fields[2], persistence: p, srcline: r.line, srcfile: r.f.Name()})
}

func newAttrDef(b *baseAttrDef) (attrDef, error) {
	c, err := getClass(b.attrType)
	if err != nil {
		return nil, fmt.Errorf("unknown attribute class %s, line %d of file %s", b.attrType, b.srcline, b.srcfile)
	}
	switch c {
	case attrClassBuiltin:
		return b, nil
	case attrClassTime:
		return &timeAttrDef{baseAttrDef: *b}, nil
	case attrClassMap:
		return newMapAttrDef(b)
	case attrClassSlice:
		return newSliceAttrDef(b)
	case attrClassStruct:
		return newStructAttrDef(b)
	}
	return nil, fmt.Errorf("Unknown attribute class %s, line %d of file %s", c, b.srcline, b.srcfile)
}

func getClass(n string) (attrClass, error) {
	switch n {
	case "byte", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "string":
		return attrClassBuiltin, nil
	case "time.Time":
		return attrClassTime, nil
	}

	switch {
	case strings.HasPrefix(n, "[]"):
		return attrClassSlice, nil
	case strings.HasPrefix(n, "map["):
		return attrClassMap, nil
	default:
		return attrClassStruct, nil
	}
}

//go:generate stringer -type=attrClass
type attrClass int

const (
	attrClassBuiltin attrClass = iota
	attrClassTime
	attrClassSlice
	attrClassMap
	attrClassStruct
)

//go:generate stringer -type=persistenceClass
type persistenceClass int

const (
	persistenceClassPersistent    persistenceClass = iota // serialized to disk and wire
	persistenceClassNonPersistent                         // serialized to wire
	persistenceClassEphemeral                             // computed or temporary storage - not serialized
)

type attrDef interface {
	AttributeID() int
	Srcline() int
	Srcfile() string
	Name() string
	Type() string
	GenerateEquals(w *writer)
	GenerateSubAttrEquals(w *writer, v1, v2 string)
}

type baseAttrDef struct {
	parentType  *typedef
	attributeID int
	name        string
	attrType    string
	persistence persistenceClass
	srcline     int
	srcfile     string
}

func (a *baseAttrDef) AttributeID() int {
	return a.attributeID
}

func (a *baseAttrDef) Srcline() int {
	return a.srcline
}

func (a *baseAttrDef) Srcfile() string {
	return a.srcfile
}

func (a *baseAttrDef) Name() string {
	return a.name
}

func (a *baseAttrDef) Type() string {
	return a.attrType
}

func (a *baseAttrDef) GenerateEquals(w *writer) {
	w.printf("  if o1.%[1]s != o2.%[1]s {\n    return false\n  }\n", a.name)
}

func (a *baseAttrDef) GenerateSubAttrEquals(w *writer, v1, v2 string) {
	w.printf("  if %s != %s {\n    return false\n  }\n", v1, v2)
}

/************************************************************************/
/************************** Time Attribute ********************/
type timeAttrDef struct {
	baseAttrDef
}

func (a *timeAttrDef) GenerateEquals(w *writer) {
	w.printf("  if !o1.%[1]s.Equal(o2.%[1]s) {\n    return false\n  }\n", a.name)
}

func (a *timeAttrDef) GenerateSubAttrEquals(w *writer, v1, v2 string) {
	w.printf("  if !%s.Equal(%s) {\n    return false\n  }\n", v1, v2)
}

/************************************************************************/
/**************************** Slice Attribute ***************************/
type sliceAttrDef struct {
	baseAttrDef
	valType string
	valAttr attrDef
}

func newSliceAttrDef(b *baseAttrDef) (*sliceAttrDef, error) {
	valType := b.attrType[2:]
	valAttr, err := newAttrDef(&baseAttrDef{attrType: valType})
	if err != nil {
		return nil, fmt.Errorf("invalid slice attribute specification %s, line %d of file %s", b.attrType, b.srcline, b.srcfile)
	}
	return &sliceAttrDef{baseAttrDef: *b, valType: valType, valAttr: valAttr}, nil
}

const sliceAttrEquals = `    if len(o1.%[1]s) != len(o2.%[1]s) {
        return false  
    }
    for idx, v1 := range o1.%[1]s {
    	v2 := o2.%[1]s[idx]
`

func (a *sliceAttrDef) GenerateEquals(w *writer) {
	w.printf(sliceAttrEquals, a.name)
	a.valAttr.GenerateSubAttrEquals(w, "v1", "v2")
	w.printf("  }\n")
}

const sliceSubAttrEquals = `    if len(%[1]s) != len(%[2]s) {
        return false  
    }
    for idx, %[1]s1 := range %[1]s {
    	%[2]s2 := %[2]s[idx]
`

func (a *sliceAttrDef) GenerateSubAttrEquals(w *writer, v1, v2 string) {
	w.printf(sliceSubAttrEquals, v1, v2)
	a.valAttr.GenerateSubAttrEquals(w, fmt.Sprintf("%s1", v1), fmt.Sprintf("%s2", v2))
	w.printf("  }\n")
}

/************************************************************************/
/**************************** Map Attribute *****************************/
type mapAttrDef struct {
	baseAttrDef
	keyType string
	keyAttr attrDef
	valType string
	valAttr attrDef
}

func newMapAttrDef(b *baseAttrDef) (*mapAttrDef, error) {
	// map[int]string
	i := strings.Index(b.attrType, "[")
	j := strings.Index(b.attrType, "]")
	if i == -1 || j == -1 {
		return nil, fmt.Errorf("invalid map attribute specification %s, line %d of file %s", b.attrType, b.srcline, b.srcfile)
	}
	keyType := b.attrType[i+1 : j]
	keyAttr, err := newAttrDef(&baseAttrDef{attrType: keyType})
	if err != nil {
		return nil, fmt.Errorf("invalid map attribute specification %s, line %d of file %s", b.attrType, b.srcline, b.srcfile)
	}
	valType := b.attrType[j+1:]
	valAttr, err := newAttrDef(&baseAttrDef{attrType: valType})
	if err != nil {
		return nil, fmt.Errorf("invalid map attribute specification %s, line %d of file %s", b.attrType, b.srcline, b.srcfile)
	}
	return &mapAttrDef{baseAttrDef: *b, keyType: keyType, keyAttr: keyAttr, valType: valType, valAttr: valAttr}, nil
}

/************************************************************************/
/**************************** Diff Obj Attribute ************************/
type structAttrDef struct {
	baseAttrDef
	packageName string
}

func newStructAttrDef(b *baseAttrDef) (*structAttrDef, error) {
	return &structAttrDef{baseAttrDef: *b}, nil
}
