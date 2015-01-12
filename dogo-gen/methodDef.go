package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func (g *Generator) parseMethod(t *typedef, fields []string) (*methodDef, error) {
	m := methodDef{parentType: t, signature: strings.Join(fields[:len(fields)-1], " "), srcline: g.r.line, srcfile: g.file.Name()}
	indentLevel := 0
	var buf bytes.Buffer
	for {
		if len(fields) == 1 && fields[0] == "}" {
			indentLevel--
		}
		for i := 0; i < indentLevel; i++ {
			buf.WriteString("    ")
		}
		buf.WriteString(strings.Join(fields, " "))
		buf.WriteString("\n")
		if fields[len(fields)-1] == "{" {
			indentLevel++
		}
		if indentLevel == 0 {
			m.body = buf.String()
			return &m, nil
		}
		var err error
		fields, err = g.r.read()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("Incomplete method specification, line %d of file %s", g.r.line, g.file.Name())
			}
			return nil, err
		}
	}
}

type methodDef struct {
	parentType *typedef
	signature  string
	body       string
	srcline    int
	srcfile    string
}
