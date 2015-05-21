//
// AUTO-GENERATED by metago. DO NOT EDIT!
//

package test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/idawes/metago"
	"time"
)

type SliceTestObject struct {
	VByte    []byte
	VUint    []uint
	VUint8   []uint8
	VUint16  []uint16
	VUint32  []uint32
	VUint64  []uint64
	VInt     []int
	VInt8    []int8
	VInt16   []int16
	VInt32   []int32
	VInt64   []int64
	VFloat32 []float32
	VFloat64 []float64
	VString  []string
	VTime    []time.Time
}

func (this *SliceTestObject) Dump() string {
	return spew.Sdump(*this)
}

func (o1 *SliceTestObject) Equals(o2 *SliceTestObject) bool {

	{
		va, vb := o1.VByte, o2.VByte
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VUint, o2.VUint
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VUint8, o2.VUint8
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VUint16, o2.VUint16
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VUint32, o2.VUint32
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VUint64, o2.VUint64
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VInt, o2.VInt
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VInt8, o2.VInt8
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VInt16, o2.VInt16
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VInt32, o2.VInt32
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VInt64, o2.VInt64
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VFloat32, o2.VFloat32
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VFloat64, o2.VFloat64
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VString, o2.VString
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if va1 != vb1 {
				return false
			}
		}
	}

	{
		va, vb := o1.VTime, o2.VTime
		if len(va) != len(vb) {
			return false
		}
		for idx, va1 := range va {
			vb1 := vb[idx]
			if !va1.Equal(vb1) {
				return false
			}
		}
	}
	return true
}

func (o1 *SliceTestObject) Diff(o2 *SliceTestObject) *metago.Diff {
	chgs := make([]metago.Chg, 0)

	{
		va, vb := o1.VByte, o2.VByte
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewByteChg(&SliceTestObjectVByteSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVByteSREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewByteChg(&SliceTestObjectVByteSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVByteSREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewByteChg(&SliceTestObjectVByteSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVByteSREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VUint, o2.VUint
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUintChg(&SliceTestObjectVUintSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUintSREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUintChg(&SliceTestObjectVUintSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUintSREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUintChg(&SliceTestObjectVUintSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUintSREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VUint8, o2.VUint8
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint8Chg(&SliceTestObjectVUint8SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint8SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint8Chg(&SliceTestObjectVUint8SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint8SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint8Chg(&SliceTestObjectVUint8SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint8SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VUint16, o2.VUint16
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint16Chg(&SliceTestObjectVUint16SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint16SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint16Chg(&SliceTestObjectVUint16SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint16SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint16Chg(&SliceTestObjectVUint16SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint16SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VUint32, o2.VUint32
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint32Chg(&SliceTestObjectVUint32SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint32SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint32Chg(&SliceTestObjectVUint32SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint32SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint32Chg(&SliceTestObjectVUint32SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint32SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VUint64, o2.VUint64
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewUint64Chg(&SliceTestObjectVUint64SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint64SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewUint64Chg(&SliceTestObjectVUint64SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint64SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewUint64Chg(&SliceTestObjectVUint64SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVUint64SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VInt, o2.VInt
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewIntChg(&SliceTestObjectVIntSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVIntSREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewIntChg(&SliceTestObjectVIntSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVIntSREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewIntChg(&SliceTestObjectVIntSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVIntSREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VInt8, o2.VInt8
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt8Chg(&SliceTestObjectVInt8SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt8SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt8Chg(&SliceTestObjectVInt8SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt8SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt8Chg(&SliceTestObjectVInt8SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt8SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VInt16, o2.VInt16
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt16Chg(&SliceTestObjectVInt16SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt16SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt16Chg(&SliceTestObjectVInt16SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt16SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt16Chg(&SliceTestObjectVInt16SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt16SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VInt32, o2.VInt32
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt32Chg(&SliceTestObjectVInt32SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt32SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt32Chg(&SliceTestObjectVInt32SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt32SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt32Chg(&SliceTestObjectVInt32SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt32SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VInt64, o2.VInt64
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewInt64Chg(&SliceTestObjectVInt64SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt64SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewInt64Chg(&SliceTestObjectVInt64SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt64SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewInt64Chg(&SliceTestObjectVInt64SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVInt64SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VFloat32, o2.VFloat32
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewFloat32Chg(&SliceTestObjectVFloat32SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVFloat32SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewFloat32Chg(&SliceTestObjectVFloat32SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVFloat32SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewFloat32Chg(&SliceTestObjectVFloat32SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVFloat32SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VFloat64, o2.VFloat64
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewFloat64Chg(&SliceTestObjectVFloat64SREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVFloat64SREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewFloat64Chg(&SliceTestObjectVFloat64SREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVFloat64SREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewFloat64Chg(&SliceTestObjectVFloat64SREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVFloat64SREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VString, o2.VString
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if va1 != vb1 {
					chgs1 = append(chgs1, metago.NewStringChg(&SliceTestObjectVStringSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVStringSREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewStringChg(&SliceTestObjectVStringSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVStringSREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewStringChg(&SliceTestObjectVStringSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVStringSREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}

	{
		va, vb := o1.VTime, o2.VTime
		for idx, va1 := range va {
			if idx < len(vb) {
				vb1 := vb[idx]
				chgs1 := make([]metago.Chg, 0)
				if !va1.Equal(vb1) {
					chgs1 = append(chgs1, metago.NewTimeChg(&SliceTestObjectVTimeSREF, vb1, va1))
				}
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVTimeSREF, idx, metago.ChangeTypeModify, chgs1))
				}
			} else {
				chgs1 := make([]metago.Chg, 0)
				chgs1 = append(chgs1, metago.NewTimeChg(&SliceTestObjectVTimeSREF, va1))
				if len(chgs1) != 0 {
					chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVTimeSREF, idx, metago.ChangeTypeDelete, chgs1))
				}
			}
		}
		for idx := len(va); idx < len(vb); idx++ {
			vb1 := vb[idx]
			chgs1 := make([]metago.Chg, 0)
			chgs1 = append(chgs1, metago.NewTimeChg(&SliceTestObjectVTimeSREF, vb1))
			if len(chgs1) != 0 {
				chgs = append(chgs, metago.NewSliceChg(&SliceTestObjectVTimeSREF, idx, metago.ChangeTypeInsert, chgs1))
			}
		}
	}
	return &metago.Diff{Chgs: chgs}
}

func (o *SliceTestObject) Apply(d *metago.Diff) error {
	for _, c := range d.Chgs {
		switch c.AttributeID() {

		case &SliceTestObjectVByteAID:
			{
				s := &o.VByte
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.ByteChg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.ByteChg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]byte, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVUintAID:
			{
				s := &o.VUint
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.UintChg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.UintChg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]uint, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVUint8AID:
			{
				s := &o.VUint8
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Uint8Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Uint8Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]uint8, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVUint16AID:
			{
				s := &o.VUint16
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Uint16Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Uint16Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]uint16, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVUint32AID:
			{
				s := &o.VUint32
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Uint32Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Uint32Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]uint32, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVUint64AID:
			{
				s := &o.VUint64
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Uint64Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Uint64Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]uint64, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVIntAID:
			{
				s := &o.VInt
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.IntChg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.IntChg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]int, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVInt8AID:
			{
				s := &o.VInt8
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Int8Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Int8Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]int8, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVInt16AID:
			{
				s := &o.VInt16
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Int16Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Int16Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]int16, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVInt32AID:
			{
				s := &o.VInt32
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Int32Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Int32Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]int32, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVInt64AID:
			{
				s := &o.VInt64
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Int64Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Int64Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]int64, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVFloat32AID:
			{
				s := &o.VFloat32
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Float32Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Float32Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]float32, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVFloat64AID:
			{
				s := &o.VFloat64
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.Float64Chg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.Float64Chg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]float64, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVStringAID:
			{
				s := &o.VString
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.StringChg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.StringChg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]string, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}

		case &SliceTestObjectVTimeAID:
			{
				s := &o.VTime
				sc := c.(*metago.SliceChg)
				idx := sc.Idx
				switch sc.Typ {
				case metago.ChangeTypeModify:
					(*s)[idx] = sc.Chgs[0].(*metago.TimeChg).NewValue
				case metago.ChangeTypeInsert:
					*s = append(*s, sc.Chgs[0].(*metago.TimeChg).NewValue)
				case metago.ChangeTypeDelete:
					new := make([]time.Time, sc.Idx)
					copy(new, *s)
					*s = new
				}
			}
		}
	}
	return nil
}
