package metago

type Diff struct {
	Chgs []Chg
}

func (d *Diff) Add(chg Chg) {
	d.Chgs = append(d.Chgs, chg)
}

type Chg interface {
	AttributeID() *AttrID
	Schemaref() *Attrdef
}

type BaseChg struct {
	schemaref *Attrdef
}

func (d *BaseChg) AttributeID() *AttrID {
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
	Chgs []Chg
}

func NewSliceChg(sref *Attrdef, idx int, typ ChangeType, chgs []Chg) Chg {
	return &SliceChg{BaseChg: BaseChg{schemaref: sref}, Idx: idx, Typ: typ, Chgs: chgs}
}

type StructChg struct {
	BaseChg
	Chg Diff
}

func NewStructChg(sref *Attrdef, chg *Diff) Chg {
	return &StructChg{BaseChg: BaseChg{schemaref: sref}, Chg: *chg}
}
