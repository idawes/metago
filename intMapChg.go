package metago

type IntMapChg struct {
    BaseChg
    Key int
    Typ ChangeType
    Chgs Diff
}

func NewIntMapChg(s *Attrdef, key int, typ ChangeType, chgs *Diff) Chg {
    return &IntMapChg{BaseChg: BaseChg{schemaref: s}, Key: key, Typ: typ, Chgs: *chgs}
}
