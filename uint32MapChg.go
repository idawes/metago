package metago

type Uint32MapChg struct {
    BaseChg
    Key uint32
    Typ ChangeType
    Chgs []Chg
}

func NewUint32MapChg(s *Attrdef, key uint32, typ ChangeType, chgs []Chg) Chg {
    return &Uint32MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
