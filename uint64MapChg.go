package metago

type Uint64MapChg struct {
    BaseChg
    Key uint64
    Typ ChangeType
    Chgs Diff
}

func NewUint64MapChg(s *Attrdef, key uint64, typ ChangeType, chgs *Diff) Chg {
    return &Uint64MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
