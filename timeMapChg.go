package metago

import "time"

type TimeMapChg struct {
	BaseChg
	Key  time.Time
	Typ  ChangeType
	Chgs []Chg
}

func NewTimeMapChg(s *Attrdef, key time.Time, typ ChangeType, chgs []Chg) Chg {
	return &TimeMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
