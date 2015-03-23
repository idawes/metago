package metago

type Uint8MapChg struct {
    BaseChg
    Key uint8
    Typ ChangeType
    Chgs []Chg
}

func NewUint8MapChg(s *Attrdef, key uint8, typ ChangeType, chgs []Chg) Chg {
    return &Uint8MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
