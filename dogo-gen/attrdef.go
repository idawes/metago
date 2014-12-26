package main

import (
	"fmt"
	"strings"
)

type attrType uint8

const (
	ATTR_CLASS_BASIC = iota
	ATTR_CLASS_TIME
	ATTR_CLASS_ARRAY
	ATTR_CLASS_SLICE
	ATTR_CLASS_MAP
	ATTR_CLASS_STRUCT
)

func (class *attrType) String() string {
	switch class {
	case ATTR_CLASS_ARRAY:
		return "Array"
	case ATTR_CLASS_BASIC:
		return "Basic"
	case ATTR_CLASS_MAP:
		return "Map"
	case ATTR_CLASS_SLICE:
		return "Slice"
	case ATTR_CLASS_STRUCT:
		return "Struct"
	case ATTR_CLASS_TIME:
		return "Time"
	}
}

type mdoAttrDef struct {
	typeId              int
	attributeId         int
	name                string
	attributeTypeShort  string
	attributeTypeFull   string
	attributeType       attrType
	collection          bool
	collectionTypeShort string
	collectionTypeFull  string
	collectionType      attrType
	persistenceType     string
	srcFileLine         int
}

func parseAttrDef(typeId, srcFileLine int, attrDefFields []string) (mdoAttrDef, error) {
	var attr mdoAttrDef
	attr.typeId = typeId
	attr.srcFileLine = srcFileLine
	var err error
	attr.attributeId, err = strconv.Atoi(attrDefFields[0])
	if err != nil {
		return nil, fmt.Errorf("Expecting an integer attribute id, found \"%s\"", attrDefFields[0])
	}

	if len(attrDefFields) < 2 {
		panic(fmt.Sprintf("Missing attribute name, line %d of file %s", typeDef.fileReader.lineIdx, typeDef.fileReader.filename))
	}
	attr.name = attrDefFields[1]

	if len(attrDefFields) < 3 {
		panic(fmt.Sprintf("Missing attribute type, line %d of file %s", typeDef.fileReader.lineIdx, typeDef.fileReader.filename))
	}
	attr.attributeTypeFull = attrDefFields[2]
	if strings.Contains(attr.attributeTypeFull, ".") {
		splitType := strings.Split(attr.attributeTypeFull, ".")
		attr.attributeTypeShort = splitType[1]
	} else {
		attr.attributeTypeShort = attr.attributeTypeFull
	}
	attr.attributeType = getAttrType(attr.attributeTypeFull, typeDef.fileReader)
	switch attr.attributeType {
	case ATTR_CLASS_SLICE:
		attr.collection = true
		attr.collectionTypeFull = attr.attributeTypeFull[2:]
		attr.collectionType = getClass(attr.collectionTypeFull, typeDef.fileReader)
	case ATTR_CLASS_ARRAY:
		attr.collection = true
		attr.collectionTypeFull = attr.attributeTypeFull[strings.Index(attr.attributeTypeShort, "]")+1:]
		attr.collectionType = getClass(attr.collectionTypeFull, typeDef.fileReader)
	}
	if strings.Contains(attr.collectionTypeFull, ".") {
		splitType := strings.Split(attr.collectionTypeFull, ".")
		attr.collectionTypeShort = splitType[1]
	} else {
		attr.collectionTypeShort = attr.collectionTypeFull
	}

	if len(attrDefFields) < 4 {
		attr.persistenceType = "persistent"
	} else {
		attr.persistenceType = attrDefFields[3]
		switch attr.persistenceType {
		case "persistent", "non-persistent", "ephemeral":
			// recognized persistance types
		default:
			panic(fmt.Sprintf("Unrecognized persistence type \"%s\", line %d of file %s", attrDefFields[3], typeDef.fileReader.lineIdx, typeDef.fileReader.filename))
		}
	}

	if len(attrDefFields) > 6 {
		panic(fmt.Sprintf("Unrecognized options \"%s\", line %d of file %s", attrDefFields[5:], typeDef.fileReader.lineIdx, typeDef.fileReader.filename))
	}

	return attr
}

func getAttrType(typeName string, fileReader *fileReader) attrType {
	switch typeName {
	case "byte", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "string":
		return ATTR_CLASS_BASIC
	case "time.Time":
		return ATTR_CLASS_TIME
	default:
		switch {
		case strings.HasPrefix(typeName, "[]"):
			return ATTR_CLASS_SLICE
		case strings.HasPrefix(typeName, "["):
			return ATTR_CLASS_ARRAY
		case strings.HasPrefix(typeName, "map"):
			return ATTR_CLASS_MAP
		default:
			return ATTR_CLASS_STRUCT
		}
	}
	panic(fmt.Sprintf("Couldn't determine attribute class for \"%s\", line %d of file %s)", typeName, fileReader.lineIdx, fileReader.filename))
}
