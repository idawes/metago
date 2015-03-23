package metago

type Int16MapChg struct {
    BaseChg
    Key int16
    Typ ChangeType
    Chgs []Chg
}

func NewInt16MapChg(s *Attrdef, key int16, typ ChangeType, chgs []Chg) Chg {
    return &Int16MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
