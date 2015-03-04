package metago

type Int64MapChg struct {
    BaseChg
    Key int64
    Typ ChangeType
    Chgs Diff
}

func NewInt64MapChg(s *Attrdef, key int64, typ ChangeType, chgs *Diff) Chg {
    return &Int64MapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
