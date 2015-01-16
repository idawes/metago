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
	c, err := getClass(fields[2])
	if err != nil {
		return nil, fmt.Errorf("Unknown attribute class %s, line %d of file %s", fields[2], g.r.line, g.file.Name())
	}
	a := baseAttrDef{parentType: t, attributeID: id, name: name, attrType: fields[2], persistence: p, srcline: g.r.line, srcfile: g.file.Name()}
	switch c {
	case attrClassBuiltin:
		return &a, nil
	case attrClassQualifiedBuiltin:
		s := strings.Split(fields[2], ".")
		a.attrType = s[1]
		return &qualifiedTypenameAttrDef{baseAttrDef: a, packageName: s[0]}, nil
	case attrClassMap:
		// map[int]string
		s := strings.Split(fields[2], "[") // [map int]string]
		if len(s) != 2 || s[0] != "map" {
			return nil, fmt.Errorf("invalid map attribute specification %s, line %d of file %s", fields[2], g.r.line, g.file.Name())
		}
		s = strings.Split(s[1], "]") // [int string]
		if len(s) != 3 {
			return nil, fmt.Errorf("invalid map attribute specification %s, line %d of file %s", fields[2], g.r.line, g.file.Name())
		}
		keyClass, err := getClass(s[1])
		if err != nil {
			return nil, fmt.Errorf("invalid map attribute specification %s, line %d of file %s", fields[2], g.r.line, g.file.Name())
		}
		switch keyClass {
		case attrClassSlice:
		case attrClassMap:
			return nil, fmt.Errorf("invalid map key class %s, line %d of file %s", keyClass, g.r.line, g.file.Name())
		}
		return &mapAttrDef{baseAttrDef: a, keyType: s[1], valType: s[2]}, nil
	case attrClassSlice:
		// []string
		s := strings.Split(fields[2], "]") // [[ string]
		if len(s) != 2 || s[0] != "[" {
			return nil, fmt.Errorf("invalid slice attribute specification %s, line %d of file %s", fields[2], g.r.line, g.file.Name())
		}
		return &sliceAttrDef{baseAttrDef: a, valType: s[1]}, nil
	case attrClassDiffableObj:
		pkg := ""
		s := strings.Split(fields[2], ".")
		if len(s) == 2 {
			pkg = s[0]
			a.attrType = s[1]
		}
		return &diffObjAttrDef{baseAttrDef: a, packageName: pkg}, nil
	}
	return nil, fmt.Errorf("Unknown attribute class %s, line %d of file %s", c, g.r.line, g.file.Name())
}

func getClass(n string) (attrClass, error) {
	switch n {
	case "byte", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "string":
		return attrClassBuiltin, nil
	case "time.Time":
		return attrClassQualifiedBuiltin, nil
	}

	switch {
	case strings.HasPrefix(n, "[]"):
		return attrClassSlice, nil
	case strings.HasPrefix(n, "map"):
		return attrClassMap, nil
	default:
		return attrClassDiffableObj, nil
	}
}

//go:generate stringer -type=attrClass
type attrClass int

const (
	attrClassBuiltin attrClass = iota
	attrClassQualifiedBuiltin
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

type qualifiedTypenameAttrDef struct {
	baseAttrDef
	packageName string
}

func (a *qualifiedTypenameAttrDef) Type() string {
	return fmt.Sprintf("%s.%s", a.packageName, a.attrType)
}

type sliceAttrDef struct {
	baseAttrDef
	valType string
}

type mapAttrDef struct {
	baseAttrDef
	keyType string
	valType string
}

type diffObjAttrDef struct {
	baseAttrDef
	packageName string
}

func (a *diffObjAttrDef) Type() string {
	return fmt.Sprintf("%s.%s", a.packageName, a.attrType)
}
