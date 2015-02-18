package metago

type ByteMapAttrChg struct {
	BaseChg
	key byte
	typ ChangeType
	chg Chg
}
