package metago

type Diff struct {
	Changes []Chg
}

type Chg interface {
	AttributeId() AttrID
	Schemaref() *Attrdef
}

type BaseChg struct {
	schemaref *Attrdef
}

func (d *BaseChg) AttributeId() AttrID {
	return d.schemaref.ID
}

func (d *BaseChg) Schemaref() *Attrdef {
	return d.schemaref
}

func (d *BaseChg) PersistenceClass() PersistenceClass {
	return d.schemaref.Persistence
}

func (d *BaseChg) WriteTo(w *Writer) error {
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

type SliceChg struct {
	BaseChg
	Idx  int
	Typ  ChangeType
	Chgs Diff
}

func NewSliceChg(sref *Attrdef, idx int, typ ChangeType, chgs *Diff) Chg {
	return &SliceChg{BaseChg: BaseChg{schemaref: sref}, Idx: idx, Typ: typ, Chgs: *chgs}
}
