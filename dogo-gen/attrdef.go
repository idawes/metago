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

func (g *Generator) parseAttribute(t *typedef, fields []string) (attrDef, error) {
	if len(fields) < 3 || len(fields) > 4 {
		return nil, fmt.Errorf("Invalid attribute specification, line %d of file %s", g.r.line, g.file.Name())
	}
	id, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("Expecting an integer attribute id, found \"%s\", line %d of file %s", fields[0], g.r.line, g.file.Name())
	}
	name := fields[1]
	p := PERSISTENCE_CLASS_PERSISTENT
	if len(fields) > 3 {
		switch fields[3] {
		case "persistent":
			p = PERSISTENCE_CLASS_PERSISTENT
		case "non-persistent":
			p = PERSISTENCE_CLASS_NON_PERSISTENT
		case "ephemeral":
			p = PERSISTENCE_CLASS_EPHEMERAL
		default:
			return nil, fmt.Errorf("Unrecognized persistence type \"%s\", line %d of file %s", fields[3], g.r.line, g.file.Name())
		}
	}
	c, err := getClass(fields[2])
	if err != nil {
		return nil, fmt.Errorf("Unknown attribute class %s, line %d of file %s", fields[2], g.r.line, g.file.Name())
	}
	a := baseAttrDef{parentType: t, attributeId: id, name: name, attrType: fields[2], persistence: p, srcline: g.r.line, srcfile: g.file.Name()}
	switch c {
	case ATTR_CLASS_BUILTIN:
		return &a, nil
	case ATTR_CLASS_QUALIFIED_BUILTIN:
		s := strings.Split(fields[2], ".")
		return &qualifiedTypenameAttrDef{baseAttrDef: a, packageName: s[0]}, nil
	case ATTR_CLASS_MAP:
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
		case ATTR_CLASS_SLICE:
		case ATTR_CLASS_MAP:
			return nil, fmt.Errorf("invalid map key class %s, line %d of file %s", keyClass, g.r.line, g.file.Name())
		}
		return &mapAttrDef{baseAttrDef: a, keyType: s[1], valType: s[2]}, nil
	case ATTR_CLASS_SLICE:
		// []string
		s := strings.Split(fields[2], "]") // [[ string]
		if len(s) != 2 || s[0] != "[" {
			return nil, fmt.Errorf("invalid slice attribute specification %s, line %d of file %s", fields[2], g.r.line, g.file.Name())
		}
		return &sliceAttrDef{baseAttrDef: a, valType: s[1]}, nil
	case ATTR_CLASS_DIFFABLE_OBJ:
		var p string
		if i := strings.LastIndex(fields[2], "."); i != -1 {
			p = fields[2][0:i]
		}
		return &diffObjAttrDef{baseAttrDef: a, packageName: p}, nil
	}
	return nil, fmt.Errorf("Unknown attribute class %s, line %d of file %s", c, g.r.line, g.file.Name())
}

func getClass(n string) (attrClass, error) {
	switch n {
	case "byte", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "string":
		return ATTR_CLASS_BUILTIN, nil
	case "time.Time":
		return ATTR_CLASS_QUALIFIED_BUILTIN, nil
	}

	switch {
	case strings.HasPrefix(n, "[]"):
		return ATTR_CLASS_SLICE, nil
	case strings.HasPrefix(n, "map"):
		return ATTR_CLASS_MAP, nil
	default:
		return ATTR_CLASS_DIFFABLE_OBJ, nil
	}
}

//go:generate stringer -type=attrClass
type attrClass int

const (
	ATTR_CLASS_BUILTIN attrClass = iota
	ATTR_CLASS_QUALIFIED_BUILTIN
	ATTR_CLASS_SLICE
	ATTR_CLASS_MAP
	ATTR_CLASS_DIFFABLE_OBJ
)

//go:generate stringer -type=persistenceClass
type persistenceClass int

const (
	PERSISTENCE_CLASS_PERSISTENT     persistenceClass = iota // serialized to disk and wire
	PERSISTENCE_CLASS_NON_PERSISTENT                         // serialized to wire
	PERSISTENCE_CLASS_EPHEMERAL                              // computed or temporary storage - not serialized
)

type attrDef interface {
	AttributeId() int
	Srcline() int
	Srcfile() string
	Imports() []string
}

type baseAttrDef struct {
	parentType  *typedef
	attributeId int
	name        string
	attrType    string
	persistence persistenceClass
	srcline     int
	srcfile     string
}

func (a *baseAttrDef) AttributeId() int {
	return a.attributeId
}

func (a *baseAttrDef) Srcline() int {
	return a.srcline
}

func (a *baseAttrDef) Srcfile() string {
	return a.srcfile
}

func (a *baseAttrDef) Imports() []string {
	return []string{}
}

type qualifiedTypenameAttrDef struct {
	baseAttrDef
	packageName string
}

func (a *qualifiedTypenameAttrDef) Imports() []string {
	return []string{a.packageName}
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

func (a *diffObjAttrDef) Imports() []string {
	if a.packageName != "" {
		return []string{a.packageName}
	}
	return []string{}
}
