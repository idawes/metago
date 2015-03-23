package metago

type ByteMapChg struct {
    BaseChg
    Key byte
    Typ ChangeType
    Chgs []Chg
}

func NewByteMapChg(s *Attrdef, key byte, typ ChangeType, chgs []Chg) Chg {
    return &ByteMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: chgs}
}
