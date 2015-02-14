package metago

import "time"

type TimeMapAttrChg struct {
	BaseAttrChg
	key time.Time
	typ ChangeType
	chg AttrChg
}
