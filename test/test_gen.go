package test

// +build: generate

//go:generate rm -f basic*.go
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
//go:generate replacer -f basic_test.tmpl -o basic_time_test.go -r "T:Time" -r "VA:time.Now()" -r "VB:time.Now().Add(5 * time.Hour)"
//go:generate goimports -w basic_time_test.go

//go:generate rm -f slice*.go
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
//go:generate replacer  -f slice_test.tmpl -o slice_time_test.go    -r "T:Time"    -r "t:time"    -r "VA:time.Now()" -r "VB:time.Now().Add(5 * time.Hour)"
//go:generate goimports -w slice_time_test.go
