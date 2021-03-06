//
// AUTO-GENERATED by metago. DO NOT EDIT!
//

package test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/idawes/metago"
	"time"
)

type BasicAttrTypesObject struct {
	VByte    byte
	VUint    uint
	VUint8   uint8
	VUint16  uint16
	VUint32  uint32
	VUint64  uint64
	VInt     int
	VInt8    int8
	VInt16   int16
	VInt32   int32
	VInt64   int64
	VFloat32 float32
	VFloat64 float64
	VString  string
	VTime    time.Time
}

func (this *BasicAttrTypesObject) ConditionalDump(t bool) string {
	if t {
		return this.Dump()
	}
	return ""
}

func (this *BasicAttrTypesObject) Dump() string {
	return spew.Sdump(*this)
}

func (o1 BasicAttrTypesObject) Equals(o2 BasicAttrTypesObject) bool {

	{
		va, vb := o1.VByte, o2.VByte
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VUint, o2.VUint
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VUint8, o2.VUint8
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VUint16, o2.VUint16
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VUint32, o2.VUint32
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VUint64, o2.VUint64
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VInt, o2.VInt
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VInt8, o2.VInt8
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VInt16, o2.VInt16
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VInt32, o2.VInt32
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VInt64, o2.VInt64
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VFloat32, o2.VFloat32
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VFloat64, o2.VFloat64
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VString, o2.VString
		if va != vb {
			return false
		}
	}

	{
		va, vb := o1.VTime, o2.VTime
		if !va.Equal(vb) {
			return false
		}
	}
	return true
}

// Diff returns a record of the differences between o1 and o2 such that applying the generated record to o1 would make it equal to o2
func (o1 BasicAttrTypesObject) Diff(o2 BasicAttrTypesObject) metago.Diff {
	chgs := make([]metago.Chg, 0)

	{
		va, vb := o1.VByte, o2.VByte
		if va != vb {
			chgs = append(chgs, metago.NewByteChg(&BasicAttrTypesObjectVByteSREF, vb, va))
		}
	}

	{
		va, vb := o1.VUint, o2.VUint
		if va != vb {
			chgs = append(chgs, metago.NewUintChg(&BasicAttrTypesObjectVUintSREF, vb, va))
		}
	}

	{
		va, vb := o1.VUint8, o2.VUint8
		if va != vb {
			chgs = append(chgs, metago.NewUint8Chg(&BasicAttrTypesObjectVUint8SREF, vb, va))
		}
	}

	{
		va, vb := o1.VUint16, o2.VUint16
		if va != vb {
			chgs = append(chgs, metago.NewUint16Chg(&BasicAttrTypesObjectVUint16SREF, vb, va))
		}
	}

	{
		va, vb := o1.VUint32, o2.VUint32
		if va != vb {
			chgs = append(chgs, metago.NewUint32Chg(&BasicAttrTypesObjectVUint32SREF, vb, va))
		}
	}

	{
		va, vb := o1.VUint64, o2.VUint64
		if va != vb {
			chgs = append(chgs, metago.NewUint64Chg(&BasicAttrTypesObjectVUint64SREF, vb, va))
		}
	}

	{
		va, vb := o1.VInt, o2.VInt
		if va != vb {
			chgs = append(chgs, metago.NewIntChg(&BasicAttrTypesObjectVIntSREF, vb, va))
		}
	}

	{
		va, vb := o1.VInt8, o2.VInt8
		if va != vb {
			chgs = append(chgs, metago.NewInt8Chg(&BasicAttrTypesObjectVInt8SREF, vb, va))
		}
	}

	{
		va, vb := o1.VInt16, o2.VInt16
		if va != vb {
			chgs = append(chgs, metago.NewInt16Chg(&BasicAttrTypesObjectVInt16SREF, vb, va))
		}
	}

	{
		va, vb := o1.VInt32, o2.VInt32
		if va != vb {
			chgs = append(chgs, metago.NewInt32Chg(&BasicAttrTypesObjectVInt32SREF, vb, va))
		}
	}

	{
		va, vb := o1.VInt64, o2.VInt64
		if va != vb {
			chgs = append(chgs, metago.NewInt64Chg(&BasicAttrTypesObjectVInt64SREF, vb, va))
		}
	}

	{
		va, vb := o1.VFloat32, o2.VFloat32
		if va != vb {
			chgs = append(chgs, metago.NewFloat32Chg(&BasicAttrTypesObjectVFloat32SREF, vb, va))
		}
	}

	{
		va, vb := o1.VFloat64, o2.VFloat64
		if va != vb {
			chgs = append(chgs, metago.NewFloat64Chg(&BasicAttrTypesObjectVFloat64SREF, vb, va))
		}
	}

	{
		va, vb := o1.VString, o2.VString
		if va != vb {
			chgs = append(chgs, metago.NewStringChg(&BasicAttrTypesObjectVStringSREF, vb, va))
		}
	}

	{
		va, vb := o1.VTime, o2.VTime
		if !va.Equal(vb) {
			chgs = append(chgs, metago.NewTimeChg(&BasicAttrTypesObjectVTimeSREF, vb, va))
		}
	}
	return metago.Diff{Chgs: chgs}
}

func (orig *BasicAttrTypesObject) Apply(d metago.Diff) error {
	for _, c := range d.Chgs {
		switch c.AttributeID() {

		case &BasicAttrTypesObjectVByteAID:
			{
				v := &orig.VByte
				*v = c.(*metago.ByteChg).NewValue
			}

		case &BasicAttrTypesObjectVUintAID:
			{
				v := &orig.VUint
				*v = c.(*metago.UintChg).NewValue
			}

		case &BasicAttrTypesObjectVUint8AID:
			{
				v := &orig.VUint8
				*v = c.(*metago.Uint8Chg).NewValue
			}

		case &BasicAttrTypesObjectVUint16AID:
			{
				v := &orig.VUint16
				*v = c.(*metago.Uint16Chg).NewValue
			}

		case &BasicAttrTypesObjectVUint32AID:
			{
				v := &orig.VUint32
				*v = c.(*metago.Uint32Chg).NewValue
			}

		case &BasicAttrTypesObjectVUint64AID:
			{
				v := &orig.VUint64
				*v = c.(*metago.Uint64Chg).NewValue
			}

		case &BasicAttrTypesObjectVIntAID:
			{
				v := &orig.VInt
				*v = c.(*metago.IntChg).NewValue
			}

		case &BasicAttrTypesObjectVInt8AID:
			{
				v := &orig.VInt8
				*v = c.(*metago.Int8Chg).NewValue
			}

		case &BasicAttrTypesObjectVInt16AID:
			{
				v := &orig.VInt16
				*v = c.(*metago.Int16Chg).NewValue
			}

		case &BasicAttrTypesObjectVInt32AID:
			{
				v := &orig.VInt32
				*v = c.(*metago.Int32Chg).NewValue
			}

		case &BasicAttrTypesObjectVInt64AID:
			{
				v := &orig.VInt64
				*v = c.(*metago.Int64Chg).NewValue
			}

		case &BasicAttrTypesObjectVFloat32AID:
			{
				v := &orig.VFloat32
				*v = c.(*metago.Float32Chg).NewValue
			}

		case &BasicAttrTypesObjectVFloat64AID:
			{
				v := &orig.VFloat64
				*v = c.(*metago.Float64Chg).NewValue
			}

		case &BasicAttrTypesObjectVStringAID:
			{
				v := &orig.VString
				*v = c.(*metago.StringChg).NewValue
			}

		case &BasicAttrTypesObjectVTimeAID:
			{
				v := &orig.VTime
				*v = c.(*metago.TimeChg).NewValue
			}
		}
	}
	return nil
}
