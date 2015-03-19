package metago

type UintMapChg struct {
    BaseChg
    Key uint
    Typ ChangeType
    Chgs Diff
}

func NewUintMapChg(s *Attrdef, key uint, typ ChangeType, chgs *Diff) Chg {
    return &UintMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
