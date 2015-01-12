package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type typedef struct {
	typeId             int
	name               string
	abstract           bool
	extendsName        string
	extends            *typedef
	definedInFileName  string
	definedOnLineNum   int
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

func (g *Generator) parseTypedef() (*typedef, error) {
	t, err := g.parseTypedefHeader()
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
func (g *Generator) parseTypedefHeader() (*typedef, error) {
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
	t := &typedef{definedInFileName: g.file.Name(), definedOnLineNum: g.r.line}
	t.typeId, err = strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("Expecting integer type number, found \"%s\", line %d of file %s", fields[1], g.r.line, g.file.Name())
	}
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
		return fmt.Errorf("Invalid attribute block header, line %d of file %s", g.r.line, g.file.Name())
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
			return fmt.Errorf("Attribute id %d redefined on line %d of file %s\n   Previous definition on line %d of file %s", a.AttributeId(), a.Srcline(), a.Srcfile(), old.Srcline(), old.Srcfile())
		}
		t.attrdefsById[a.AttributeId()] = a
		for _, i := range a.Imports() {
			t.imports[i] = struct{}{}
		}
		if *veryVerbose {
			fmt.Printf("      Found attribute spec: %v, Imports: %v\n", a, a.Imports())
		}
	}
}

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
		return fmt.Errorf("Invalid method block header, line %d of file %s", g.r.line, g.file.Name())
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
			t.imports[strings.Replace(fields[1], "\"", "", -1)] = struct{}{}
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
