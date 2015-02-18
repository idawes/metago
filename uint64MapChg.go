package metago

type Uint64MapAttrChg struct {
	BaseChg
	key uint64
	typ ChangeType
	chg Chg
}
