package metago

type StringMapChg struct {
    BaseChg
    Key string
    Typ ChangeType
    Chgs []Chg
}

func NewStringMapChg(s *Attrdef, key string, typ ChangeType, chgs []Chg) Chg {
    return &StringMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
