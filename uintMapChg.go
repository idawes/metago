package metago

type UintMapChg struct {
    BaseChg
    Key uint
    Typ ChangeType
    Chgs []Chg
}

func NewUintMapChg(s *Attrdef, key uint, typ ChangeType, chgs []Chg) Chg {
    return &UintMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
