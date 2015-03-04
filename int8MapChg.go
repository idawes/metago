package metago

type Int8MapChg struct {
    BaseChg
    Key int8
    Typ ChangeType
    Chgs Diff
}

func NewInt8MapChg(s *Attrdef, key int8, typ ChangeType, chgs *Diff) Chg {
    return &Int8MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
