//
// AUTO-GENERATED by metago. DO NOT EDIT!
//

package test

import (
	"github.com/idawes/metago"
	"time"
)

type MapTestObject struct {
	VByteByte   map[byte]byte
	VByteUint   map[byte]uint
	VByteUint8  map[byte]uint8
	VByteUint16 map[byte]uint16
	VByteUint32 map[byte]uint32
	VByteUint64 map[byte]uint64
	VByteInt    map[byte]int
	VByteInt8   map[byte]int8
	VByteInt16  map[byte]int16
	VByteInt32  map[byte]int32
	VByteInt64  map[byte]int64
	VByteTime   map[byte]time.Time
}

func (o1 MapTestObject) Equals(o2 MapTestObject) bool {

	{
		va, vb := o1.VByteByte, o2.VByteByte
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteUint, o2.VByteUint
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteUint8, o2.VByteUint8
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteUint16, o2.VByteUint16
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteUint32, o2.VByteUint32
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteUint64, o2.VByteUint64
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteInt, o2.VByteInt
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteInt8, o2.VByteInt8
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteInt16, o2.VByteInt16
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteInt32, o2.VByteInt32
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteInt64, o2.VByteInt64
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if va1 != vb1 {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}

	{
		va, vb := o1.VByteTime, o2.VByteTime
		if len(va) != len(vb) {
			return false
		}
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				if !va1.Equal(vb1) {
					return false
				}
			} else {
				return false // didn't find key in vb
			}
		}
	}
	return true
}

// Diff returns a record of the differences between o1 and o2 such that applying the generated record to o1 would make it equal to o2
func (o1 MapTestObject) Diff(o2 MapTestObject) metago.Diff {
	chgs := make([]metago.Chg, 0)

	{
		va, vb := o1.VByteByte, o2.VByteByte
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewByteChg(&MapTestObjectVByteByteSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteByteSREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewByteChg(&MapTestObjectVByteByteSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteByteSREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewByteChg(&MapTestObjectVByteByteSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteByteSREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteUint, o2.VByteUint
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUintChg(&MapTestObjectVByteUintSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUintSREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUintChg(&MapTestObjectVByteUintSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUintSREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUintChg(&MapTestObjectVByteUintSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUintSREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteUint8, o2.VByteUint8
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint8Chg(&MapTestObjectVByteUint8SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint8SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint8Chg(&MapTestObjectVByteUint8SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint8SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint8Chg(&MapTestObjectVByteUint8SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint8SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteUint16, o2.VByteUint16
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint16Chg(&MapTestObjectVByteUint16SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint16SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint16Chg(&MapTestObjectVByteUint16SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint16SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint16Chg(&MapTestObjectVByteUint16SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint16SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteUint32, o2.VByteUint32
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint32Chg(&MapTestObjectVByteUint32SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint32SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint32Chg(&MapTestObjectVByteUint32SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint32SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint32Chg(&MapTestObjectVByteUint32SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint32SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteUint64, o2.VByteUint64
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint64Chg(&MapTestObjectVByteUint64SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint64SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint64Chg(&MapTestObjectVByteUint64SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint64SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint64Chg(&MapTestObjectVByteUint64SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteUint64SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteInt, o2.VByteInt
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewIntChg(&MapTestObjectVByteIntSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteIntSREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewIntChg(&MapTestObjectVByteIntSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteIntSREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewIntChg(&MapTestObjectVByteIntSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteIntSREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteInt8, o2.VByteInt8
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt8Chg(&MapTestObjectVByteInt8SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt8SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt8Chg(&MapTestObjectVByteInt8SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt8SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt8Chg(&MapTestObjectVByteInt8SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt8SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteInt16, o2.VByteInt16
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt16Chg(&MapTestObjectVByteInt16SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt16SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt16Chg(&MapTestObjectVByteInt16SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt16SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt16Chg(&MapTestObjectVByteInt16SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt16SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteInt32, o2.VByteInt32
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt32Chg(&MapTestObjectVByteInt32SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt32SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt32Chg(&MapTestObjectVByteInt32SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt32SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt32Chg(&MapTestObjectVByteInt32SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt32SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteInt64, o2.VByteInt64
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt64Chg(&MapTestObjectVByteInt64SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt64SREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt64Chg(&MapTestObjectVByteInt64SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt64SREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt64Chg(&MapTestObjectVByteInt64SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteInt64SREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VByteTime, o2.VByteTime
		for key, va1 := range va {
			if vb1, ok := vb[key]; ok {
				// "key" exists in both "va" and "vb"
				chgs1 := make([]metago.Chg, 0)
				if !va1.Equal(vb1) {
					chgs1 = append(chgs1, metago.NewTimeChg(&MapTestObjectVByteTimeSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteTimeSREF, key, metago.ChangeTypeModify, chgs1))
				}
			} else {
				// "key" exists in "va" but not in "vb"
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewTimeChg(&MapTestObjectVByteTimeSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteTimeSREF, key, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for key, vb1 := range vb {
			if _, ok := va[key]; ok {
				continue
			}
			// "key" exists in vb but not int va"
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewTimeChg(&MapTestObjectVByteTimeSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewByteMapChg(&MapTestObjectVByteTimeSREF, key, metago.ChangeTypeInsert, chgs1))
			}
		}
	}
	return metago.Diff{Chgs: chgs}
}

func (orig *MapTestObject) Apply(d metago.Diff) error {
	for _, c := range d.Chgs {
		switch c.AttributeID() {

		case &MapTestObjectVByteByteAID:
			{
				m := &orig.VByteByte
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.ByteChg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.ByteChg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteUintAID:
			{
				m := &orig.VByteUint
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.UintChg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.UintChg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteUint8AID:
			{
				m := &orig.VByteUint8
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Uint8Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Uint8Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteUint16AID:
			{
				m := &orig.VByteUint16
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Uint16Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Uint16Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteUint32AID:
			{
				m := &orig.VByteUint32
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Uint32Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Uint32Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteUint64AID:
			{
				m := &orig.VByteUint64
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Uint64Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Uint64Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteIntAID:
			{
				m := &orig.VByteInt
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.IntChg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.IntChg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteInt8AID:
			{
				m := &orig.VByteInt8
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Int8Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Int8Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteInt16AID:
			{
				m := &orig.VByteInt16
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Int16Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Int16Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteInt32AID:
			{
				m := &orig.VByteInt32
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Int32Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Int32Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteInt64AID:
			{
				m := &orig.VByteInt64
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					m[key] = mc.Chgs[0].(*metago.Int64Chg).NewValue
				case metago.ChangeTypeInsert:
					m[key] = mc.Chgs[0].(*metago.Int64Chg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}

		case &MapTestObjectVByteTimeAID:
			{
				m := &orig.VByteTime
				mc := c.(*metago.ByteMapChg)
				key := mc.Key
				switch mc.Typ {
				case metago.ChangeTypeModify:
					*m[key] = mc.Chgs[0].(*metago.TimeChg).NewValue
				case metago.ChangeTypeInsert:
					*m[key] = mc.Chgs[0].(*metago.TimeChg).NewValue
				case metago.ChangeTypeDelete:
					delete(*m, key)
					if len(*m) == 0 {
						*m = nil
					}
				}
			}
		}
	}
	return nil
}
