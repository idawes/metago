package metago

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

type Attrdef struct {
	ID          AttrID
	Persistence PersistenceClass
}

func ProcessDiff(d Diff) {
	// for _, c := range d.Changes() {
	// if c.AttributeId() == test.SliceTestObjectByteSliceId {
	// }
	// }
}
