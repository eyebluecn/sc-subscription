package util

import (
	"reflect"
	"time"
)

// Any is a helper routine that allocates a new interface value
// to store v and returns a pointer to it.
//
//	// Usage: var _ *Type = pointer.Any(Type(value) | value).(*Type)
//
//	var _ *bool = pointer.Any(true).(*bool)
//	var _ *byte = pointer.Any(byte(1)).(*byte)
//	var _ *complex64 = pointer.Any(complex64(1.1)).(*complex64)
//	var _ *complex128 = pointer.Any(complex128(1.1)).(*complex128)
//	var _ *float32 = pointer.Any(float32(1.1)).(*float32)
//	var _ *float64 = pointer.Any(float64(1.1)).(*float64)
//	var _ *int = pointer.Any(int(1)).(*int)
//	var _ *int8 = pointer.Any(int8(8)).(*int8)
//	var _ *int16 = pointer.Any(int16(16)).(*int16)
//	var _ *int32 = pointer.Any(int32(32)).(*int32)
//	var _ *int64 = pointer.Any(int64(64)).(*int64)
//	var _ *rune = pointer.Any(rune(1)).(*rune)
//	var _ *string = pointer.Any("ptr").(*string)
//	var _ *uint = pointer.Any(uint(1)).(*uint)
//	var _ *uint8 = pointer.Any(uint8(8)).(*uint8)
//	var _ *uint16 = pointer.Any(uint16(16)).(*uint16)
//	var _ *uint32 = pointer.Any(uint32(32)).(*uint32)
//	var _ *uint64 = pointer.Any(uint64(64)).(*uint64)
//	var _ *uintptr = pointer.Any(uintptr(64)).(*uintptr)
func Any(v interface{}) interface{} {
	r := reflect.New(reflect.TypeOf(v))
	reflect.ValueOf(r.Interface()).Elem().Set(reflect.ValueOf(v))
	return r.Interface()
}

// BoolPointer is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func BoolPointer(v bool) *bool { return &v }

// BytePointer is a helper routine that allocates a new byte value
// to store v and returns a pointer to it.
func BytePointer(v byte) *byte { return &v }

// Complex64Pointer is a helper routine that allocates a new complex64 value
// to store v and returns a pointer to it.
func Complex64Pointer(v complex64) *complex64 { return &v }

// Complex128Pointer is a helper routine that allocates a new complex128 value
// to store v and returns a pointer to it.
func Complex128Pointer(v complex128) *complex128 { return &v }

// Float32Pointer is a helper routine that allocates a new float32 value
// to store v and returns a pointer to it.
func Float32Pointer(v float32) *float32 { return &v }

// Float64Pointer is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64Pointer(v float64) *float64 { return &v }

// IntPointer is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func IntPointer(v int) *int { return &v }

// Int8Pointer is a helper routine that allocates a new int8 value
// to store v and returns a pointer to it.
func Int8Pointer(v int8) *int8 { return &v }

// Int16Pointer is a helper routine that allocates a new int16 value
// to store v and returns a pointer to it.
func Int16Pointer(v int16) *int16 { return &v }

// Int32Pointer is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it.
func Int32Pointer(v int32) *int32 { return &v }

// Int64Pointer is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64Pointer(v int64) *int64 { return &v }

// RunePointer is a helper routine that allocates a new rune value
// to store v and returns a pointer to it.
func RunePointer(v rune) *rune { return &v }

// StringPointer is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func StringPointer(v string) *string { return &v }

// UintPointer is a helper routine that allocates a new uint value
// to store v and returns a pointer to it.
func UintPointer(v uint) *uint { return &v }

// Uint8Pointer is a helper routine that allocates a new uint8 value
// to store v and returns a pointer to it.
func Uint8Pointer(v uint8) *uint8 { return &v }

// Uint16Pointer is a helper routine that allocates a new uint16 value
// to store v and returns a pointer to it.
func Uint16Pointer(v uint16) *uint16 { return &v }

// Uint32Pointer is a helper routine that allocates a new uint32 value
// to store v and returns a pointer to it.
func Uint32Pointer(v uint32) *uint32 { return &v }

// Uint64Pointer is a helper routine that allocates a new uint64 value
// to store v and returns a pointer to it.
func Uint64Pointer(v uint64) *uint64 { return &v }

// TimePointer is a helper routine that allocates a new time.Time value
// to store v and returns a pointer to it.
func TimePointer(v time.Time) *time.Time { return &v }

// DurationPointer is a helper routine that allocates a new time.Duration value
// to store v and returns a pointer to it.
func DurationPointer(v time.Duration) *time.Duration { return &v }
