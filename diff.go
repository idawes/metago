package metago

type Diff interface {
	Changes() []AttrDiff
}

type AttrDiff interface {
	AttributeId() AttrID
}

type BaseAttrDiff struct {
	schemaref *Attrdef
}

func (d *BaseAttrDiff) AttributeId() AttrID {
	return d.schemaref.ID
}

func (d *BaseAttrDiff) PersistenceClass() PersistenceClass {
	return d.schemaref.Persistence
}

func (d *BaseAttrDiff) WriteTo(w *Writer) error {
	w.Write(d.schemaref.ID.pkg[:])
	w.WriteVarint(int64(d.schemaref.ID.typ))
	w.WriteVarint(int64(d.schemaref.ID.attr))
	return nil
}
