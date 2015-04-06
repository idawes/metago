package metago

import "fmt"

type Typedef struct {
	ID TypeID
}

//go:generate stringer -type=PersistenceClass
type PersistenceClass int

const (
	PersistenceClassPersistent    PersistenceClass = iota // serialized to disk and wire
	PersistenceClassNonPersistent                         // serialized to wire
	PersistenceClassEphemeral                             // computed or temporary storage - not serialized
)

func (p PersistenceClass) ShortString() string {
	switch p {
	case PersistenceClassPersistent:
		return "(P)"
	case PersistenceClassNonPersistent:
		return "(N)"
	case PersistenceClassEphemeral:
		return "(E)"
	}
	return "(?)"
}

type Attrdef struct {
	ID          *AttrID
	Persistence PersistenceClass
}

func (a *Attrdef) String() string {
	return fmt.Sprintf("%s.%s.%s - %s", a.ID.PkgName, a.ID.TypeID.Name, a.ID.Name, a.Persistence.ShortString())
}
