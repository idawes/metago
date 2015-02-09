// Copyright 2015 Ian Dawes. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"container/list"
	"os"
	"strings"
)

type reader struct {
	filename    string
	f           *os.File
	r           *bufio.Reader
	line        int
	unreadLines *list.List
	err         error
}

func newReader(filename string) *reader {
	f, err := os.Open(filename)
	if err != nil {
		return &reader{err: err}
	}
	return &reader{filename: filename, f: f, r: bufio.NewReader(f), unreadLines: list.New()}
}

func (r *reader) close() {
	r.f.Close() // ok to ignore errors here, because f is read-only and this is an ephemeral executable
}

func (r *reader) read() ([]string, error) {
	var (
		line string
		err  error
	)
	if e := r.unreadLines.Front(); e != nil {
		r.unreadLines.Remove(e)
		return e.Value.([]string), nil
	}
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

func (r *reader) unread(fields []string) {
	r.unreadLines.PushFront(fields)
}
