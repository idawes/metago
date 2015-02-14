package metago

// +build: generate

//go:generate replacer -f generic_diff.tmpl -o byteDiff.go     -r "T:Byte"   -r "t:byte"   -r "s:1+1"
//go:generate replacer -f generic_diff.tmpl -o uint8Diff.go    -r "T:Uint8"  -r "t:uint8"  -r "s:1+1"
//go:generate replacer -f generic_diff.tmpl -o uint16Diff.go   -r "T:Uint16" -r "t:uint16" -r "s:2+2"
//go:generate replacer -f generic_diff.tmpl -o uint32Diff.go   -r "T:Uint32" -r "t:uint32" -r "s:4+1"
//go:generate replacer -f generic_diff.tmpl -o uint64Diff.go   -r "T:Uint64" -r "t:uint64" -r "s:8+1"
//go:generate replacer -f generic_diff.tmpl -o int8Diff.go     -r "T:Int8"   -r "t:int8"   -r "s:1+1"
//go:generate replacer -f generic_diff.tmpl -o int16Diff.go    -r "T:Int16"  -r "t:int16"  -r "s:2+2"
//go:generate replacer -f generic_diff.tmpl -o int32Diff.go    -r "T:Int32"  -r "t:int32"  -r "s:4+4"
//go:generate replacer -f generic_diff.tmpl -o int64Diff.go    -r "T:Int64"  -r "t:int64"  -r "s:8+8"
//go:generate replacer -f generic_diff.tmpl -o stringDiff.go    -r "T:String"  -r "t:string"  -r "s:uint32(4+len(d.OldValue)+4+len(d.NewValue))"
//go:generate replacer -f generic_diff.tmpl -o timeDiff.go    -r "T:Time"  -r "t:time.Time"
//go:generate goimports -w timeDiff.go

//go:generate replacer -f generic_mapChg.tmpl -o byteMapChg.go     -r "T:Byte"   -r "t:byte"
//go:generate replacer -f generic_mapChg.tmpl -o uint8MapChg.go    -r "T:Uint8"  -r "t:uint8"
//go:generate replacer -f generic_mapChg.tmpl -o uint16MapChg.go   -r "T:Uint16" -r "t:uint16"
//go:generate replacer -f generic_mapChg.tmpl -o uint32MapChg.go   -r "T:Uint32" -r "t:uint32"
//go:generate replacer -f generic_mapChg.tmpl -o uint64MapChg.go   -r "T:Uint64" -r "t:uint64"
//go:generate replacer -f generic_mapChg.tmpl -o int8MapChg.go     -r "T:Int8"   -r "t:int8"
//go:generate replacer -f generic_mapChg.tmpl -o int16MapChg.go    -r "T:Int16"  -r "t:int16"
//go:generate replacer -f generic_mapChg.tmpl -o int32MapChg.go    -r "T:Int32"  -r "t:int32"
//go:generate replacer -f generic_mapChg.tmpl -o int64MapChg.go    -r "T:Int64"  -r "t:int64"
//go:generate replacer -f generic_mapChg.tmpl -o stringMapChg.go    -r "T:String"  -r "t:string"
//go:generate replacer -f generic_mapChg.tmpl -o timeMapChg.go    -r "T:Time"  -r "t:time.Time"
//go:generate goimports -w timeMapChg.go
