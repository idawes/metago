package metago

import "time"

type TimeMapAttrChg struct {
	BaseChg
	key time.Time
	typ ChangeType
	chg Chg
}
