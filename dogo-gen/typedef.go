package main

type typedef struct {
	typeId             int
	name               string
	abstract           bool
	extendsName        string
	extends            *typedef
	definedInFileName  string
	definedOnLineNum   int
	attributeImports   []string
	attrdefsById       map[int]attrDef
	attrsdefsInIdOrder attrdefList
	methodImports      []string
	abstractMethods    []string
	methods            []string
	imports            map[string]string
	persistent         bool
}

type attrdefList []attrDef

func (l attrdefList) Len() int           { return len(l) }
func (l attrdefList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l attrdefList) Less(i, j int) bool { return l[i].attributeId < l[j].attributeId }
