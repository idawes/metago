package metago

type Uint64MapAttrChg struct {
    BaseAttrChg
    key uint64
    typ ChangeType
    chg AttrChg
}
