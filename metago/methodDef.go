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
func parseMethod(t *typedef, r *reader) (*methodDef, error) {
	fields, _ := r.read()
	sig := strings.Join(fields[1:], " ")
	i := strings.Index(sig, "(")
	if i == -1 {
		return nil, fmt.Errorf("invalid method signature (missing \"(\" for parameters) in %s, line %d of file %s", fields[1], r.line, r.f.Name())
	}

	j := strings.Index(sig, ")")
	if j == -1 {
		return nil, fmt.Errorf("invalid method signature (missing \")\" for parameters) in %s, line %d of file %s", fields[1], r.line, r.f.Name())
	}
	if strings.Index(sig, "{") != len(sig)-1 {
		return nil, fmt.Errorf("invalid method signature (missing \")\" for parameters) in %s, line %d of file %s", fields[1], r.line, r.f.Name())
	}
	m := methodDef{parentType: t, name: sig[:i], params: sig[i+1 : j], returns: strings.TrimSpace(sig[j+1 : len(sig)-1]), srcline: r.line, srcfile: r.f.Name()}
	indentLevel := 1
	var buf bytes.Buffer
	for {
		var err error
		fields, err = r.read()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("incomplete method specification, line %d of file %s", r.line, r.f.Name())
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
