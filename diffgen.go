package metago

// +build: generate

//go:generate replacer -f genericBaseChg.tmpl -o byteChg.go     -r "T:Byte"    -r "t:byte"
//go:generate replacer -f genericBaseChg.tmpl -o uintChg.go     -r "T:Uint"    -r "t:uint"
//go:generate replacer -f genericBaseChg.tmpl -o uint8Chg.go    -r "T:Uint8"   -r "t:uint8"
//go:generate replacer -f genericBaseChg.tmpl -o uint16Chg.go   -r "T:Uint16"  -r "t:uint16"
//go:generate replacer -f genericBaseChg.tmpl -o uint32Chg.go   -r "T:Uint32"  -r "t:uint32"
//go:generate replacer -f genericBaseChg.tmpl -o uint64Chg.go   -r "T:Uint64"  -r "t:uint64"
//go:generate replacer -f genericBaseChg.tmpl -o intChg.go      -r "T:Int"     -r "t:int"
//go:generate replacer -f genericBaseChg.tmpl -o int8Chg.go     -r "T:Int8"    -r "t:int8"
//go:generate replacer -f genericBaseChg.tmpl -o int16Chg.go    -r "T:Int16"   -r "t:int16"
//go:generate replacer -f genericBaseChg.tmpl -o int32Chg.go    -r "T:Int32"   -r "t:int32"
//go:generate replacer -f genericBaseChg.tmpl -o int64Chg.go    -r "T:Int64"   -r "t:int64"
//go:generate replacer -f genericBaseChg.tmpl -o stringChg.go   -r "T:String"  -r "t:string"
//go:generate replacer -f genericBaseChg.tmpl -o float32Chg.go  -r "T:Float32" -r "t:float32"
//go:generate replacer -f genericBaseChg.tmpl -o float64Chg.go  -r "T:Float64" -r "t:float64"
//go:generate replacer -f genericBaseChg.tmpl -o timeChg.go     -r "T:Time"    -r "t:time.Time"
//go:generate goimports -w timeChg.go

//go:generate replacer -f genericMapChg.tmpl -o byteMapChg.go     -r "T:Byte"   -r "t:byte"
//go:generate replacer -f genericMapChg.tmpl -o uintMapChg.go     -r "T:Uint"   -r "t:uint"
//go:generate replacer -f genericMapChg.tmpl -o uint8MapChg.go    -r "T:Uint8"  -r "t:uint8"
//go:generate replacer -f genericMapChg.tmpl -o uint16MapChg.go   -r "T:Uint16" -r "t:uint16"
//go:generate replacer -f genericMapChg.tmpl -o uint32MapChg.go   -r "T:Uint32" -r "t:uint32"
//go:generate replacer -f genericMapChg.tmpl -o uint64MapChg.go   -r "T:Uint64" -r "t:uint64"
//go:generate replacer -f genericMapChg.tmpl -o intMapChg.go      -r "T:Int"    -r "t:int"
//go:generate replacer -f genericMapChg.tmpl -o int8MapChg.go     -r "T:Int8"   -r "t:int8"
//go:generate replacer -f genericMapChg.tmpl -o int16MapChg.go    -r "T:Int16"  -r "t:int16"
//go:generate replacer -f genericMapChg.tmpl -o int32MapChg.go    -r "T:Int32"  -r "t:int32"
//go:generate replacer -f genericMapChg.tmpl -o int64MapChg.go    -r "T:Int64"  -r "t:int64"
//go:generate replacer -f genericMapChg.tmpl -o stringMapChg.go   -r "T:String" -r "t:string"
//go:generate replacer -f genericMapChg.tmpl -o timeMapChg.go     -r "T:Time"   -r "t:time.Time"
//go:generate goimports -w timeMapChg.go
