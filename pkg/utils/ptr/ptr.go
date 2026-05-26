/*
Package ptr provides utility functions for converting non-addressable primitive types to pointers.
Its useful in contexts where a variable gives nil primitive type pointers semantics
(often meaning "not set") which can make it annoying to set the value.

Example

	type Foo struct {
		A *int
	}

	func main() {
		foo := Foo{
			A: ptr.Int(1)
		}
	}
*/
package ptr

func Int(v int) *int { _ = "STUB: not implemented"; return nil }

func Int8(v int8) *int8 { _ = "STUB: not implemented"; return nil }

func Int16(v int16) *int16 { _ = "STUB: not implemented"; return nil }

func Int32(v int32) *int32 { _ = "STUB: not implemented"; return nil }

func Int64(v int64) *int64 { _ = "STUB: not implemented"; return nil }

func Uint(v uint) *uint { _ = "STUB: not implemented"; return nil }

func Uint8(v uint8) *uint8 { _ = "STUB: not implemented"; return nil }

func Uint16(v uint16) *uint16 { _ = "STUB: not implemented"; return nil }

func Uint32(v uint32) *uint32 { _ = "STUB: not implemented"; return nil }

func Uint64(v uint64) *uint64 { _ = "STUB: not implemented"; return nil }

func Float32(v float32) *float32 { _ = "STUB: not implemented"; return nil }

func Float64(v float64) *float64 { _ = "STUB: not implemented"; return nil }

func String(v string) *string { _ = "STUB: not implemented"; return nil }

func Bool(v bool) *bool { _ = "STUB: not implemented"; return nil }

func Byte(v byte) *byte { _ = "STUB: not implemented"; return nil }

func Rune(v rune) *rune { _ = "STUB: not implemented"; return nil }

func Complex64(v complex64) *complex64 { _ = "STUB: not implemented"; return nil }

func Complex128(v complex128) *complex128 { _ = "STUB: not implemented"; return nil }
