package main

import (
	"bufio"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"sort"
	"strconv"
	"strings"
)

type typeID struct {
	uuid  *uuid.UUID
	subID int
}

func (t *typeID) String() string {
	return fmt.Sprintf("%s:%d", t.uuid, t.subID)
}

type typedef struct {
	typeID              typeID
	name                string
	abstract            bool
	extendsName         string
	extends             *typedef
	srcfile             string
	srcline             int
	attrDefsByID        map[int]attrDef
	attrDefsByIDInOrder attrdefList
	abstractMethods     []string
	methods             []*methodDef
	imports             map[string]struct{}
	persistent          bool
}

type attrdefList []attrDef

func (l attrdefList) Len() int           { return len(l) }
func (l attrdefList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l attrdefList) Less(i, j int) bool { return l[i].AttributeID() < l[j].AttributeID() }

/*
	Main internal entry point for parsing a type definition.

	type <typeID> <typeName> <concrete|abstract> [extends <typeName>] {
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
func (g *generator) parseTypedef(uuid *uuid.UUID) (*typedef, error) {
	t, err := g.parseTypedefHeader(uuid)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}
	if err := g.parseImportBlock(t); err != nil {
		return nil, err
	}
	if err := g.parseAttributeBlock(t); err != nil {
		return nil, err
	}
	if err := g.parseMethodBlock(t); err != nil {
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
		type <typeID> <typeName> <concrete|abstract> [extends <typeName>] {
*/
func (g *generator) parseTypedefHeader(uuid *uuid.UUID) (*typedef, error) {
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
	subID, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("Expecting integer type number, found \"%s\", line %d of file %s", fields[1], g.r.line, g.file.Name())
	}
	t.typeID = typeID{uuid: uuid, subID: subID}
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

func (g *generator) parseImportBlock(t *typedef) error {
	t.imports = make(map[string]struct{})
	fields, err := g.r.read()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("Incomplete type specification, line %d of file %s", g.r.line, g.file.Name())
		}
		return err
	}
	if len(fields) != 2 || fields[0] != "imports" || fields[1] != "{" {
		if *veryVerbose {
			fmt.Println("      No imports specified")
		}
		g.r.unread(fields)
		return nil
	}
	for {
		fields, err := g.r.read()
		if err != nil {
			return err
		}
		if len(fields) != 1 {
			return fmt.Errorf("Invalid import block entry, line %d of file %s", g.r.line, g.file.Name())
		}
		if fields[0] == "}" { // naked "}" means end of block.
			if *veryVerbose {
				fmt.Printf("      Found imports: %v\n", t.imports)
			}
			return nil
		}
		t.imports[strings.Replace(fields[0], "\"", "", -1)] = struct{}{}
	}
}

/*
	attribute block:
		attributes {
			<attributedef>*
		}
*/
func (g *generator) parseAttributeBlock(t *typedef) error {
	t.attrDefsByID = make(map[int]attrDef, 10)
	fields, err := g.r.read()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("Incomplete type specification, line %d of file %s", g.r.line, g.file.Name())
		}
		return err
	}
	if len(fields) != 2 || fields[0] != "attributes" || fields[1] != "{" {
		if *veryVerbose {
			fmt.Println("      No attributes specified")
		}
		g.r.unread(fields)
		return nil
	}
	for {
		fields, err := g.r.read()
		if err != nil {
			return err
		}
		if len(fields) == 1 && fields[0] == "}" { // naked "}" means end of block.
			// generate sorted attribute list before returning
			if len(t.attrDefsByID) > 0 {
				t.attrDefsByIDInOrder = make([]attrDef, 0)
				for _, a := range t.attrDefsByID {
					t.attrDefsByIDInOrder = append(t.attrDefsByIDInOrder, a)
				}
				sort.Sort(t.attrDefsByIDInOrder)
			}
			return nil
		}
		a, err := g.parseAttribute(t, fields)
		if err != nil {
			return err
		}
		if old, present := t.attrDefsByID[a.AttributeID()]; present {
			return fmt.Errorf("Duplicate defintion of attribute id %d on line %d of file %s\n   It is also defined on line %d of file %s", a.AttributeID(), a.Srcline(), a.Srcfile(), old.Srcline(), old.Srcfile())
		}
		t.attrDefsByID[a.AttributeID()] = a
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
func (g *generator) parseMethodBlock(t *typedef) error {
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
			fmt.Println("      No methods specified")
		}
		g.r.unread(fields)
		return nil
	}
	for {
		fields, err := g.r.read()
		if err != nil {
			return err
		}
		if len(fields) == 1 && fields[0] == "}" { // naked "}" means end of block.
			return nil
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

func (t *typedef) validateTypeHierarchy(typesByName map[string]*typedef) error {
	if t.extendsName == "" {
		return nil
	}
	super, ok := typesByName[t.extendsName]
	if !ok {
		return fmt.Errorf("Couldn't find type %s while validating hierarchy for type %s, defined on line %d of file %s", t.extendsName, t.name, t.srcline, t.srcfile)
	}
	t.extends = super
	return nil
}

func (t *typedef) generate(w *bufio.Writer) error {
	if _, err := w.WriteString("import (\n"); err != nil {
		return err
	}
	if err := t.generateImports(w); err != nil {
		return err
	}
	if _, err := w.WriteString(fmt.Sprintf(")\n\ntype %s struct {\n", t.name)); err != nil {
		return err
	}
	if err := t.generateAttributes(w); err != nil {
		return err
	}
	if _, err := w.WriteString("}\n"); err != nil {
		return err
	}
	return nil
}

func (t *typedef) generateImports(w *bufio.Writer) error {
	t.coalesceImports(t.imports)
	for im, _ := range t.imports {
		if _, err := w.WriteString(fmt.Sprintf("  %s\n", im)); err != nil {
			return err
		}
	}
	return nil
}

func (t *typedef) coalesceImports(imports map[string]struct{}) {
	if t.extends != nil {
		t.extends.coalesceImports(imports)
	}
	for k, v := range t.imports {
		imports[k] = v
	}
}

func (t *typedef) generateAttributes(w *bufio.Writer) error {
	if t.extends != nil {
		t.extends.generateAttributes(w)
	}
	for _, a := range t.attrDefsByIDInOrder {
		if _, err := w.WriteString(fmt.Sprintf("  %s %s\n", a.Name(), a.Type())); err != nil {
			return err
		}
	}
	return nil
}
