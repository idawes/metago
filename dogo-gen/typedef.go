package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"strconv"
	"strings"
)

type typeId struct {
	uuid  *uuid.UUID
	subid int
}

func (t *typeId) String() string {
	return fmt.Sprintf("%s:%d", t.uuid, t.subid)
}

type typedef struct {
	typeId             typeId
	name               string
	abstract           bool
	extendsName        string
	extends            *typedef
	srcfile            string
	srcline            int
	attrdefsById       map[int]attrDef
	attrsdefsInIdOrder attrdefList
	abstractMethods    []string
	methods            []*methodDef
	imports            map[string]struct{}
	persistent         bool
}

type attrdefList []attrDef

func (l attrdefList) Len() int           { return len(l) }
func (l attrdefList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l attrdefList) Less(i, j int) bool { return l[i].AttributeId() < l[j].AttributeId() }

/*
	Main internal entry point for parsing a type definition.

	type <typeId> <typeName> <concrete|abstract> [extends <typeName>] {
		attributes {
			<attributedef>*
		}
		methods {
			<import|method>*
		}

	If this type extends another type, the supertype must be defined in the same package. Any method defined on a subtype that has the same name as a method in the supertype must also
	have the same signature.

	Imports may occur at any point in the method block, allowing them to be specified just above the method that requires it. Imports may be duplicated, they will only be present once in the generated
	output.
*/
func (g *Generator) parseTypedef(uuid *uuid.UUID) (*typedef, error) {
	t, err := g.parseTypedefHeader(uuid)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}
	t.imports = make(map[string]struct{})
	err = g.parseAttributeBlock(t)
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("Incomplete type specification, line %d of file %s", g.r.line, g.file.Name())
		}
		return nil, err
	}
	err = g.parseMethodBlock(t)
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("Incomplete type specification, line %d of file %s", g.r.line, g.file.Name())
		}
		return nil, err
	}
	fields, err := g.r.read()
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("Incomplete type specification, line %d of file %s", g.r.line, g.file.Name())
		}
		return nil, err
	}
	if len(fields) != 1 && fields[0] != "}" {
		return nil, fmt.Errorf("Missing closing \"}\", line %d of file %s", g.r.line, g.file.Name())
	}
	return t, nil
}

/*
	typedef header:
		type <typeId> <typeName> <concrete|abstract> [extends <typeName>] {
*/
func (g *Generator) parseTypedefHeader(uuid *uuid.UUID) (*typedef, error) {
	fields, err := g.r.read()
	if err != nil {
		return nil, err
	}
	if *veryVerbose {
		fmt.Printf("    Parsing %q as type def header from line %d of file %s\n", fields, g.r.line, g.file.Name())
	}
	if len(fields) < 4 {
		return nil, fmt.Errorf("Invalid type specification, line %d of file %s", g.r.line, g.file.Name())
	}
	if fields[0] != "type" {
		return nil, fmt.Errorf("Expecting \"type\", found \"%s\", line %d o file %s", fields[0], g.r.line, g.file.Name())
	}
	t := &typedef{srcfile: g.file.Name(), srcline: g.r.line}
	subId, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("Expecting integer type number, found \"%s\", line %d of file %s", fields[1], g.r.line, g.file.Name())
	}
	t.typeId = typeId{uuid: uuid, subid: subId}
	t.name = fields[2]
	if fields[3] == "abstract" {
		t.abstract = true
	} else if fields[3] == "concrete" {
		t.abstract = false
	} else {
		return nil, fmt.Errorf("Expecting \"abstract\" or \"concrete\", found \"%s\", line %d of file %s", fields[3], g.r.line, g.file.Name())
	}
	if fields[len(fields)-1] != "{" {
		return nil, fmt.Errorf("Missing \"{\", line %d of file %s", g.r.line, g.file.Name())
	}
	if len(fields) == 5 {
		return t, nil
	}
	if fields[4] != "extends" {
		return nil, fmt.Errorf("Expecting \"extends\", found \"%s\", line %d of file %s", fields[5], g.r.line, g.file.Name())
	}
	if len(fields) < 7 {
		return nil, fmt.Errorf("Missing base type name, line %d of file %s", g.r.line, g.file.Name())
	}
	t.extendsName = fields[5]

	return t, nil
}

/*
	attribute block:
		attributes {
			<attributedef>*
		}
*/
func (g *Generator) parseAttributeBlock(t *typedef) error {
	t.attrdefsById = make(map[int]attrDef, 10)
	fields, err := g.r.read()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("Incomplete type specification, line %d of file %s", g.r.line, g.file.Name())
		}
		return err
	}
	if len(fields) != 2 || fields[0] != "attributes" || fields[1] != "{" {
		if *veryVerbose {
			fmt.Println("    No attributes specified")
		}
		g.r.unread(fields)
		return nil
	}
	for {
		fields, err := g.r.read()
		if err != nil {
			return err
		}
		if len(fields) == 1 && fields[0] == "}" { // naked "}" means end of attribute block.
			return nil
		}
		a, err := g.parseAttribute(t, fields)
		if err != nil {
			return err
		}
		if old, present := t.attrdefsById[a.AttributeId()]; present {
			return fmt.Errorf("Duplicate defintion of attribute id %d on line %d of file %s\n   It is also defined on line %d of file %s", a.AttributeId(), a.Srcline(), a.Srcfile(), old.Srcline(), old.Srcfile())
		}
		t.attrdefsById[a.AttributeId()] = a
		for _, i := range a.Imports() {
			t.imports[i] = struct{}{}
		}
		if *veryVerbose {
			fmt.Printf("      Found attribute spec: %v\n", a)
		}
	}
}

/*
	method block:
		methods {
			<import|method>*
		}

	import:
		standard go import spec.

	method:
		standard golang method decl, except leaving out the receiver spec, which will be automatically generated. Special syntax ##super##(<args>) can be used to invoke function of same name in supertype.
*/
func (g *Generator) parseMethodBlock(t *typedef) error {
	t.methods = make([]*methodDef, 0)
	fields, err := g.r.read()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("Incomplete type specification, line %d of file %s", g.r.line, g.file.Name())
		}
		return err
	}
	if len(fields) != 2 || fields[0] != "methods" || fields[1] != "{" {
		if *veryVerbose {
			fmt.Println("    No methods specified")
		}
		g.r.unread(fields)
		return nil
	}
	for {
		fields, err := g.r.read()
		if err != nil {
			return err
		}
		if len(fields) == 1 && fields[0] == "}" { // naked "}" means end of method block.
			return nil
		}
		if len(fields) == 2 && fields[0] == "import" {
			t.imports[strings.Join(fields, " ")] = struct{}{}
			continue
		}
		if len(fields) > 0 && fields[0] == "func" {
			m, err := g.parseMethod(t, fields[1:])
			if err != nil {
				return err
			}
			t.methods = append(t.methods, m)
			continue
		}
		return fmt.Errorf("Invalid keyword: \"%s\", line %d of file %s", fields[0], g.r.line, g.file.Name())
	}
}
