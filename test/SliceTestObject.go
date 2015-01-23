//
// AUTO-GENERATED by metago. DO NOT EDIT!
//

package test

import (
	"github.com/davecgh/go-spew/spew"
	"time"
)

type SliceTestObject struct {
	ByteSlice     []byte
	U8Slice       []uint8
	U16Slice      []uint16
	U32Slice      []uint32
	U64Slice      []uint64
	S8Slice       []int8
	S16Slice      []int16
	S32Slice      []int32
	S64Slice      []int64
	StringSlice   []string
	TimeSlice     []time.Time
	String2dSlice [][]string
}

func (this *SliceTestObject) Dump() string {
	return spew.Sdump(*this)
}

func (o1 *SliceTestObject) Equals(other interface{}) bool {
	switch o2 := other.(type) {
	case *SliceTestObject:
		return o1.equals(o2)
	case SliceTestObject:
		return o1.equals(&o2)
	}
	return false
}

func (o1 *SliceTestObject) equals(o2 *SliceTestObject) bool {

	//---------  comparison for ByteSlice ----------------------------------/
	if len(o1.ByteSlice) != len(o2.ByteSlice) {
		return false
	}
	for idx, v1 := range o1.ByteSlice {
		v2 := o2.ByteSlice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for U8Slice ----------------------------------/
	if len(o1.U8Slice) != len(o2.U8Slice) {
		return false
	}
	for idx, v1 := range o1.U8Slice {
		v2 := o2.U8Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for U16Slice ----------------------------------/
	if len(o1.U16Slice) != len(o2.U16Slice) {
		return false
	}
	for idx, v1 := range o1.U16Slice {
		v2 := o2.U16Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for U32Slice ----------------------------------/
	if len(o1.U32Slice) != len(o2.U32Slice) {
		return false
	}
	for idx, v1 := range o1.U32Slice {
		v2 := o2.U32Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for U64Slice ----------------------------------/
	if len(o1.U64Slice) != len(o2.U64Slice) {
		return false
	}
	for idx, v1 := range o1.U64Slice {
		v2 := o2.U64Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for S8Slice ----------------------------------/
	if len(o1.S8Slice) != len(o2.S8Slice) {
		return false
	}
	for idx, v1 := range o1.S8Slice {
		v2 := o2.S8Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for S16Slice ----------------------------------/
	if len(o1.S16Slice) != len(o2.S16Slice) {
		return false
	}
	for idx, v1 := range o1.S16Slice {
		v2 := o2.S16Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for S32Slice ----------------------------------/
	if len(o1.S32Slice) != len(o2.S32Slice) {
		return false
	}
	for idx, v1 := range o1.S32Slice {
		v2 := o2.S32Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for S64Slice ----------------------------------/
	if len(o1.S64Slice) != len(o2.S64Slice) {
		return false
	}
	for idx, v1 := range o1.S64Slice {
		v2 := o2.S64Slice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for StringSlice ----------------------------------/
	if len(o1.StringSlice) != len(o2.StringSlice) {
		return false
	}
	for idx, v1 := range o1.StringSlice {
		v2 := o2.StringSlice[idx]
		if v1 != v2 {
			return false
		}
	}

	//---------  comparison for TimeSlice ----------------------------------/
	if len(o1.TimeSlice) != len(o2.TimeSlice) {
		return false
	}
	for idx, v1 := range o1.TimeSlice {
		v2 := o2.TimeSlice[idx]
		if !v1.Equal(v2) {
			return false
		}
	}

	//---------  comparison for String2dSlice ----------------------------------/
	if len(o1.String2dSlice) != len(o2.String2dSlice) {
		return false
	}
	for idx, v1 := range o1.String2dSlice {
		v2 := o2.String2dSlice[idx]
		if len(v1) != len(v2) {
			return false
		}
		for idx, v11 := range v1 {
			v22 := v2[idx]
			if v11 != v22 {
				return false
			}
		}
	}
	return true
}
