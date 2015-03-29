package test

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestBasic(t *testing.T) {
	a := StructTestObject{}
	b := StructTestObject{}

	d := a.Diff(&b)
	if len(d.Chgs) != 0 {
		t.Errorf("Found %d differences between a and b, expected 0", len(d.Chgs))
		spew.Dump(d)
	}

	oa := StructTestObject{MB: make(map[int]BasicAttrTypesObject, 0)}
	oa.MB[0] = BasicAttrTypesObject{S8Field: 22}
	oa.B.TimeField = time.Now()
	ob := StructTestObject{MB: make(map[int]BasicAttrTypesObject, 0), B: BasicAttrTypesObject{ByteField: 3}}
	ob.B.TimeField = oa.B.TimeField.Add(time.Minute)
	ob.MB[0] = BasicAttrTypesObject{S8Field: 55}
	ob.MB[1] = BasicAttrTypesObject{S8Field: 56}
	ob.MB[2] = BasicAttrTypesObject{S8Field: 57}
	ob.MB[3] = BasicAttrTypesObject{S8Field: 58}
	ob.MB[4] = BasicAttrTypesObject{S8Field: 59}
	ob.MB[5] = BasicAttrTypesObject{S8Field: 60}
	d = oa.Diff(&ob)
	if len(d.Chgs) != 7 {
		t.Errorf("Found %d differences between a and b, expected 1", len(d.Chgs))
		spew.Dump(d)
	}

	oa.Apply(d)
	if !oa.Equals(&ob) {
		t.Errorf("Objects aren't equal")
		spew.Dump(d)
	}
}

func BenchmarkBasic(b *testing.B) {
	oa := StructTestObject{MB: make(map[int]BasicAttrTypesObject, 0)}
	ob := StructTestObject{MB: make(map[int]BasicAttrTypesObject, 0), B: BasicAttrTypesObject{ByteField: 3}}
	ob.MB[0] = BasicAttrTypesObject{S8Field: 55}
	ob.MB[1] = BasicAttrTypesObject{S8Field: 56}
	ob.MB[2] = BasicAttrTypesObject{S8Field: 57}
	ob.MB[3] = BasicAttrTypesObject{S8Field: 58}
	ob.MB[4] = BasicAttrTypesObject{S8Field: 59}
	ob.MB[5] = BasicAttrTypesObject{S8Field: 60}
	for i := 0; i < b.N; i++ {
		oa.Diff(&ob)
	}
}
