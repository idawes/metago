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
	attrdefsById       map[int]attrdef
	attrsdefsInIdOrder attrdefList
	methodImports      []string
	abstractMethods    []string
	methods            []string
	imports            map[string]string
	persistent         bool
}
