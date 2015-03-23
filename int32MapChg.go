package metago

type Int32MapChg struct {
    BaseChg
    Key int32
    Typ ChangeType
    Chgs []Chg
}

func NewInt32MapChg(s *Attrdef, key int32, typ ChangeType, chgs []Chg) Chg {
    return &Int32MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
