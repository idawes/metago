package metago

import "time"

type TimeMapChg struct {
	BaseChg
	Key  time.Time
	Typ  ChangeType
	Chgs Diff
}

func NewTimeMapChg(s *Attrdef, key time.Time, typ ChangeType, chgs *Diff) Chg {
	return &TimeMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
