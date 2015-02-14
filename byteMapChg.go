package metago

type ByteMapAttrChg struct {
    BaseAttrChg
    key byte
    typ ChangeType
    chg AttrChg
}
