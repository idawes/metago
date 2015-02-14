package metago

type Diff struct {
	Changes []AttrChg
}

type AttrChg interface {
	AttributeId() AttrID
}

type BaseAttrChg struct {
	schemaref *Attrdef
}

func (d *BaseAttrChg) AttributeId() AttrID {
	return d.schemaref.ID
}

func (d *BaseAttrChg) PersistenceClass() PersistenceClass {
	return d.schemaref.Persistence
}

func (d *BaseAttrChg) WriteTo(w *Writer) error {
	w.Write(d.schemaref.ID.Pkg[:])
	w.WriteVarint(int64(d.schemaref.ID.Typ))
	w.WriteVarint(int64(d.schemaref.ID.Attr))
	return nil
}

//go:generate stringer -type=ChangeType
type ChangeType int

const (
	ChangeTypeInsert ChangeType = iota
	ChangeTypeDelete
	ChangeTypeModify
)

type SliceAttrChg struct {
	BaseAttrChg
	idx int
	typ ChangeType
	chg AttrChg
}
