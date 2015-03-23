package metago

type Uint16MapChg struct {
    BaseChg
    Key uint16
    Typ ChangeType
    Chgs []Chg
}

func NewUint16MapChg(s *Attrdef, key uint16, typ ChangeType, chgs []Chg) Chg {
    return &Uint16MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
