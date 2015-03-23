package metago

type IntMapChg struct {
    BaseChg
    Key int
    Typ ChangeType
    Chgs []Chg
}

func NewIntMapChg(s *Attrdef, key int, typ ChangeType, chgs []Chg) Chg {
    return &IntMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
