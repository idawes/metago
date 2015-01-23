//
// AUTO-GENERATED by metago. DO NOT EDIT!
//

package test

import (
	"github.com/davecgh/go-spew/spew"
	"time"
)

type BasicAttrTypesObject struct {
	ByteField   byte
	U8Field     uint8
	U16Field    uint16
	U32Field    uint32
	U64Field    uint64
	S8Field     int8
	S16Field    int16
	S32Field    int32
	S64Field    int64
	StringField string
	TimeField   time.Time
}

func (this *BasicAttrTypesObject) Dump() (string, error) {
	return spew.Sdump(*this)
}

func (this *BasicAttrTypesObject) ConditionalDump(t bool) string {
	if t {
		return this.Dump()
	}
	return ""
}

func (o1 *BasicAttrTypesObject) Equals(other interface{}) bool {
	switch o2 := other.(type) {
	case *BasicAttrTypesObject:
		return o1.equals(o2)
	case BasicAttrTypesObject:
		return o1.equals(&o2)
	}
	return false
}

func (o1 *BasicAttrTypesObject) equals(o2 *BasicAttrTypesObject) bool {

	//---------  comparison for ByteField ----------------------------------/
	if o1.ByteField != o2.ByteField {
		return false
	}

	//---------  comparison for U8Field ----------------------------------/
	if o1.U8Field != o2.U8Field {
		return false
	}

	//---------  comparison for U16Field ----------------------------------/
	if o1.U16Field != o2.U16Field {
		return false
	}

	//---------  comparison for U32Field ----------------------------------/
	if o1.U32Field != o2.U32Field {
		return false
	}

	//---------  comparison for U64Field ----------------------------------/
	if o1.U64Field != o2.U64Field {
		return false
	}

	//---------  comparison for S8Field ----------------------------------/
	if o1.S8Field != o2.S8Field {
		return false
	}

	//---------  comparison for S16Field ----------------------------------/
	if o1.S16Field != o2.S16Field {
		return false
	}

	//---------  comparison for S32Field ----------------------------------/
	if o1.S32Field != o2.S32Field {
		return false
	}

	//---------  comparison for S64Field ----------------------------------/
	if o1.S64Field != o2.S64Field {
		return false
	}

	//---------  comparison for StringField ----------------------------------/
	if o1.StringField != o2.StringField {
		return false
	}

	//---------  comparison for TimeField ----------------------------------/
	if !o1.TimeField.Equal(o2.TimeField) {
		return false
	}
	return true
}