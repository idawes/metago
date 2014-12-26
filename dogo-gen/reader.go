package main

import (
	"bufio"
	"io"
	"strings"
)

type Reader struct {
	SkipCommentLines bool
	r                *bufio.Reader
	line             int
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: bufio.NewReader(r)}
}

func (r *Reader) Read() ([]string, error) {
	var line string
	inBlockComment := false
	for {
		r.line++
		line, err := r.r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if r.SkipCommentLines {
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
		}
		break
	}
	return strings.Fields(line), nil
}
