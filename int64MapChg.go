package metago

type Int64MapChg struct {
    BaseChg
    Key int64
    Typ ChangeType
    Chgs []Chg
}

func NewInt64MapChg(s *Attrdef, key int64, typ ChangeType, chgs []Chg) Chg {
    return &Int64MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
