package metago

type StringMapAttrChg struct {
	BaseChg
	key string
	typ ChangeType
	chg Chg
}
