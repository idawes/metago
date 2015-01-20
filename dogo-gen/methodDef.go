package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

/*
	Method Spec:
	 func <name>(<params>) [(]<returns>[)] {
	 	<body>
	 }
*/
func (g *generator) parseMethod(t *typedef, fields []string) (*methodDef, error) {
	sig := strings.Join(fields[1:], " ")
	i := strings.Index(sig, "(")
	if i == -1 {
		return nil, fmt.Errorf("Invalid method signature (missing \"(\" for parameters) in %s, line %d of file %s", fields[1], g.r.line, g.file.Name())
	}

	j := strings.Index(sig, ")")
	if j == -1 {
		return nil, fmt.Errorf("Invalid method signature (missing \")\" for parameters) in %s, line %d of file %s", fields[1], g.r.line, g.file.Name())
	}
	if strings.Index(sig, "{") != len(sig)-1 {
		return nil, fmt.Errorf("Invalid method signature (missing \")\" for parameters) in %s, line %d of file %s", fields[1], g.r.line, g.file.Name())
	}
	m := methodDef{parentType: t, name: sig[:i], params: sig[i+1 : j], returns: strings.TrimSpace(sig[j+1 : len(sig)-1]), srcline: g.r.line, srcfile: g.file.Name()}
	indentLevel := 1
	var buf bytes.Buffer
	for {
		var err error
		fields, err = g.r.read()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("Incomplete method specification, line %d of file %s", g.r.line, g.file.Name())
			}
			return nil, err
		}
		if len(fields) == 1 && fields[0] == "}" {
			indentLevel--
		}
		if indentLevel == 0 {
			m.body = buf.String()
			return &m, nil
		}
		for i := 0; i < indentLevel; i++ {
			buf.WriteString("    ")
		}
		buf.WriteString(strings.Join(fields, " "))
		buf.WriteString("\n")
		if fields[len(fields)-1] == "{" {
			indentLevel++
		}
	}
}

type methodDef struct {
	parentType *typedef
	name       string
	params     string
	returns    string
	body       string
	srcline    int
	srcfile    string
}
