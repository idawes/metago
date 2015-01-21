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

func (g *generator) parseAttribute(t *typedef, fields []string) (attrDef, error) {
	if len(fields) < 3 || len(fields) > 4 {
		return nil, fmt.Errorf("Invalid attribute specification, line %d of file %s", g.r.line, g.file.Name())
	}
	id, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("Expecting an integer attribute id, found \"%s\", line %d of file %s", fields[0], g.r.line, g.file.Name())
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
			return nil, fmt.Errorf("Unrecognized persistence type \"%s\", line %d of file %s", fields[3], g.r.line, g.file.Name())
		}
	}
	return newAttrDef(&baseAttrDef{parentType: t, attributeID: id, name: name, attrType: fields[2], persistence: p, srcline: g.r.line, srcfile: g.file.Name()})
}

func newAttrDef(b *baseAttrDef) (attrDef, error) {
	c, err := getClass(b.attrType)
	if err != nil {
		return nil, fmt.Errorf("Unknown attribute class %s, line %d of file %s", b.attrType, b.srcline, b.srcfile)
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
	case attrClassDiffableObj:
		return newDiffableObjAttrDef(b)
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
		return attrClassDiffableObj, nil
	}
}

//go:generate stringer -type=attrClass
type attrClass int

const (
	attrClassBuiltin attrClass = iota
	attrClassTime
	attrClassSlice
	attrClassMap
	attrClassDiffableObj
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
	GenerateEquals(g *generator)
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

func (a *baseAttrDef) GenerateEquals(g *generator) {
	g.printf("  if o1.%[1]s != o2.%[1]s {\n    return false\n  }\n", a.name)
}

/************************************************************************/
/************************** Qualified Type Attribute ********************/
type timeAttrDef struct {
	baseAttrDef
}

func (a *timeAttrDef) GenerateEquals(g *generator) {
	g.printf("  if !o1.%[1]s.Equals(o2.%[1]s) {\n    return false\n  }\n", a.name)
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

const sliceEquals = `    if len(o1.%[1]s) != len(o2.%[1]s) {
        return false  
    }
    for idx, newVal := range o1.%[1]s {
`

func (a *sliceAttrDef) GenerateEquals(g *generator) {
	g.printf(sliceEquals, a.name)

	g.printf("}\n")
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
type diffObjAttrDef struct {
	baseAttrDef
	packageName string
}

func newDiffableObjAttrDef(b *baseAttrDef) (*diffObjAttrDef, error) {
	return &diffObjAttrDef{baseAttrDef: *b}, nil
}
