package metago

import (
	"bufio"
	"io"
)

type Writer struct {
	*bufio.Writer
	p PersistenceClass
}

func NewWriter(w io.Writer, p PersistenceClass) *Writer {
	return &Writer{Writer: bufio.NewWriter(w), p: p}
}

type Writable interface {
	WriteTo(w *Writer) error
	PersistenceClass() PersistenceClass
}

func (w *Writer) WriteObj(obj Writable) error {
	if w.p == PersistenceClassPersistent && obj.PersistenceClass() != PersistenceClassPersistent {
		return nil
	}
	if w.p == PersistenceClassNonPersistent && obj.PersistenceClass() == PersistenceClassEphemeral {
		return nil
	}
	return obj.WriteTo(w)
}

func (w *Writer) WriteUVarint(v uint64) int {
	i := 0
	for v >= 0x80 {
		w.WriteByte(byte(v) | 0x80)
		v >>= 7
		i++
	}
	w.WriteByte(byte(v))
	return i + 1
}

func (w *Writer) WriteVarint(v int64) int {
	uv := uint64(v) << 1
	if v < 0 {
		uv = ^uv
	}
	return w.WriteUVarint(uv)
}

func (w *Writer) WriteBool(v bool) int {
	if v {
		w.WriteByte(1)
	} else {
		w.WriteByte(0)
	}
	return 1
}

func (w *Writer) WriteString(v string) int {
	c, _ := w.Write([]byte(v))
	return c
}
