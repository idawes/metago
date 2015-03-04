package metago

type Int32MapChg struct {
    BaseChg
    Key int32
    Typ ChangeType
    Chgs Diff
}

func NewInt32MapChg(s *Attrdef, key int32, typ ChangeType, chgs *Diff) Chg {
    return &Int32MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
