package main

import (
	"fmt"
	"strconv"
	"strings"
)

//go:generate stringer -type=attrClass
type attrClass uint8

const (
	ATTR_TYPE_BASIC attrClass = iota
	ATTR_TYPE_TIME
	ATTR_TYPE_SLICE
	ATTR_TYPE_MAP
	ATTR_TYPE_STRUCT
)

type attrDef struct {
	typeId              int
	attributeId         int
	name                string
	attributeTypeShort  string
	attributeTypeFull   string
	class               attrClass
	collection          bool
	collectionTypeShort string
	collectionTypeFull  string
	collectionType      attrClass
	persistenceType     string
	srcline             int
}

func newAttrDef(typeId, srcline int, attrDefFields []string) (*attrDef, error) {
	a := attrDef{typeId: typeId, srcline: srcline}

	var err error
	a.attributeId, err = strconv.Atoi(attrDefFields[0])
	if err != nil {
		return nil, fmt.Errorf("Expecting an integer attribute id, found \"%s\"", attrDefFields[0])
	}

	if len(attrDefFields) < 2 {
		return nil, fmt.Errorf("Missing attribute name")
	}
	a.name = attrDefFields[1]

	if len(attrDefFields) < 3 {
		return nil, fmt.Errorf("Missing attribute type")
	}
	a.attributeTypeFull = attrDefFields[2]
	if strings.Contains(a.attributeTypeFull, ".") {
		splitType := strings.Split(a.attributeTypeFull, ".")
		a.attributeTypeShort = splitType[1]
	} else {
		a.attributeTypeShort = a.attributeTypeFull
	}

	a.class, err = getClass(a.attributeTypeFull)
	if err != nil {

	}
	switch a.class {
	case ATTR_TYPE_SLICE:
		a.collection = true
		a.collectionTypeFull = a.attributeTypeFull[2:]
		a.collectionType, err = getClass(a.collectionTypeFull)
	case ATTR_TYPE_MAP:
		a.collection = true
	}
	if strings.Contains(a.collectionTypeFull, ".") {
		splitType := strings.Split(a.collectionTypeFull, ".")
		a.collectionTypeShort = splitType[1]
	} else {
		a.collectionTypeShort = a.collectionTypeFull
	}

	if len(attrDefFields) < 4 {
		a.persistenceType = "persistent"
	} else {
		a.persistenceType = attrDefFields[3]
		switch a.persistenceType {
		case "persistent", "non-persistent", "ephemeral":
			// recognized persistance types
		default:
			return nil, fmt.Errorf("Unrecognized persistence type \"%s\"", attrDefFields[3])
		}
	}

	if len(attrDefFields) > 6 {
		return nil, fmt.Errorf("Unrecognized options \"%s\"", attrDefFields[5:])
	}

	return &a, nil
}

func getClass(typeName string) (attrClass, error) {
	switch typeName {
	case "byte", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "string":
		return ATTR_TYPE_BASIC, nil
	case "time.Time":
		return ATTR_TYPE_TIME, nil
	default:
		switch {
		case strings.HasPrefix(typeName, "[]"):
			return ATTR_TYPE_SLICE, nil
		case strings.HasPrefix(typeName, "map"):
			return ATTR_TYPE_MAP, nil
		default:
			return ATTR_TYPE_STRUCT, nil
		}
	}
	panic(fmt.Sprintf("Couldn't determine attribute class for \"%s\")", typeName))
}
