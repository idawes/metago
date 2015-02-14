package metago

type StringMapAttrChg struct {
    BaseAttrChg
    key string
    typ ChangeType
    chg AttrChg
}
