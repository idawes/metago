package main

import (
	"bufio"
	"io"
	"strings"
)

type reader struct {
	r    *bufio.Reader
	line int
}

func newReader(r io.Reader) *reader {
	return &reader{r: bufio.NewReader(r)}
}

func (r *reader) read() ([]string, error) {
	var (
		line string
		err  error
	)
	inBlockComment := false
	for {
		r.line++
		line, err = r.r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "//") {
			continue
		}
		if inBlockComment {
			if strings.HasSuffix(line, "*/") {
				inBlockComment = false
			}
			continue
		}
		if strings.HasPrefix(line, "/*") {
			inBlockComment = true
			continue
		}
		break
	}
	return strings.Fields(line), nil
}
