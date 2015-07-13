package test

// +build: generate

//go:generate rm -f "`ls -1 *test.go`"

//go:generate go install github.com/idawes/metago/metago
//go:generate go install github.com/idawes/replacer
//go:generate metago github.com/idawes/metago/test

//go:generate replacer -f basic_test.tmpl -o basic_byte_test.go    -r "T:Byte"    -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_uint_test.go    -r "T:Uint"    -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_uint8_test.go   -r "T:Uint8"   -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_uint16_test.go  -r "T:Uint16"  -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_uint32_test.go  -r "T:Uint32"  -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_uint64_test.go  -r "T:Uint64"  -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_int_test.go     -r "T:Int"     -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_int8_test.go    -r "T:Int8"    -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_int16_test.go   -r "T:Int16"   -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_int32_test.go   -r "T:Int32"   -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_int64_test.go   -r "T:Int64"   -r "VA:3" -r "VB:5"
//go:generate replacer -f basic_test.tmpl -o basic_float32_test.go -r "T:Float32" -r "VA:3.34" -r "VB:5.42"
//go:generate replacer -f basic_test.tmpl -o basic_float64_test.go -r "T:Float64" -r "VA:3.23" -r "VB:5.332"
//go:generate replacer -f basic_test.tmpl -o basic_string_test.go -r "T:String" -r "VA:\"Foo\"" -r "VB:\"Bar\""
//go:generate replacer -f basic_test.tmpl -o basic_time_test.go -r "T:Time" -r "VA:time.Unix(1436000000, 0)" -r "VB:time.Unix(1436100000, 0)"
//go:generate goimports -w basic_time_test.go

//go:generate replacer  -f slice_test.tmpl -o slice_byte_test.go    -r "T:Byte"    -r "t:byte"    -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_uint_test.go    -r "T:Uint"    -r "t:uint"    -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_uint8_test.go   -r "T:Uint8"   -r "t:uint8"   -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_uint16_test.go  -r "T:Uint16"  -r "t:uint16"  -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_uint32_test.go  -r "T:Uint32"  -r "t:uint32"  -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_uint64_test.go  -r "T:Uint64"  -r "t:uint64"  -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_int_test.go     -r "T:Int"     -r "t:int"     -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_int8_test.go    -r "T:Int8"    -r "t:int8"    -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_int16_test.go   -r "T:Int16"   -r "t:int16"   -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_int32_test.go   -r "T:Int32"   -r "t:int32"   -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_int64_test.go   -r "T:Int64"   -r "t:int64"   -r "VA:3"          -r "VB:5"
//go:generate replacer  -f slice_test.tmpl -o slice_float32_test.go -r "T:Float32" -r "t:float32" -r "VA:3.34"       -r "VB:5.42"
//go:generate replacer  -f slice_test.tmpl -o slice_float64_test.go -r "T:Float64" -r "t:float64" -r "VA:3.23"       -r "VB:5.332"
//go:generate replacer  -f slice_test.tmpl -o slice_string_test.go  -r "T:String"  -r "t:string"  -r "VA:\"Foo\""    -r "VB:\"Bar\""
//go:generate replacer  -f slice_test.tmpl -o slice_time_test.go    -r "T:Time"    -r "t:time"    -r "VA:time.Unix(1436000000, 0)" -r "VB:time.Unix(1436100000, 0)"
//go:generate goimports -w slice_time_test.go

//go:generate replacer -f map_test.tmpl -o map_byte_byte_test.go   -r "T:ByteByte"   -r "kt:byte" -r "vt:byte"   -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_uint_test.go   -r "T:ByteUint"   -r "kt:byte" -r "vt:uint"   -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_uint8_test.go  -r "T:ByteUint8"  -r "kt:byte" -r "vt:uint8"  -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_uint16_test.go -r "T:ByteUint16" -r "kt:byte" -r "vt:uint16" -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_uint32_test.go -r "T:ByteUint32" -r "kt:byte" -r "vt:uint32" -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_uint64_test.go -r "T:ByteUint64" -r "kt:byte" -r "vt:uint64" -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_int_test.go    -r "T:ByteInt"    -r "kt:byte" -r "vt:int"    -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_int8_test.go   -r "T:ByteInt8"   -r "kt:byte" -r "vt:int8"   -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_int16_test.go  -r "T:ByteInt16"  -r "kt:byte" -r "vt:int16"  -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_int32_test.go  -r "T:ByteInt32"  -r "kt:byte" -r "vt:int32"  -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
//go:generate replacer -f map_test.tmpl -o map_byte_int64_test.go  -r "T:ByteInt64"  -r "kt:byte" -r "vt:int64"  -r "K:2" -r "KI:3" -r "VA:3" -r "VB:5"
// //go:generate replacer  -f map_test.tmpl -o map_float32_test.go -r "T:Float32" -r "t:float32" -r "VA:3.34"       -r "VB:5.42"
// //go:generate replacer  -f map_test.tmpl -o map_float64_test.go -r "T:Float64" -r "t:float64" -r "VA:3.23"       -r "VB:5.332"
// //go:generate replacer  -f map_test.tmpl -o map_string_test.go  -r "T:String"  -r "t:string"  -r "VA:\"Foo\""    -r "VB:\"Bar\""
// //go:generate replacer  -f map_test.tmpl -o map_time_test.go    -r "T:Time"    -r "t:time"    -r "VA:time.Unix(1436000000, 0)" -r "VB:time.Unix(1436100000, 0)"
// //go:generate goimports -w map_time_test.go
