package metago

type StringMapChg struct {
    BaseChg
    Key string
    Typ ChangeType
    Chgs Diff
}

func NewStringMapChg(s *Attrdef, key string, typ ChangeType, chgs *Diff) Chg {
    return &StringMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
