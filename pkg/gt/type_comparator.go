package gt

import "reflect"

// TODO: add struct type comparators

func OfBool[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Bool
}

func OfInt[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Int
}

func OfInt8[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Int8
}

func OfInt16[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Int16
}

func OfInt32[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Int32
}

func OfInt64[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Int64
}

func OfUint[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Uint
}

func OfUint8[A any](actual *A) bool {
	return reflect.TypeOf(actual).Kind() == reflect.Uint8
}

func OfUint16[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Uint16
}

func OfUint32[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Uint32
}

func OfUint64[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Uint64
}

func OfUintptr[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Uintptr
}

func OfFloat32[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Float32
}

func OfFloat64[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Float64
}

func OfString[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.String
}

func OfArray[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Array
}

func OfChan[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Chan
}

func OfFunc[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Func
}

func OfInterface[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Interface
}

func OfMap[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Map
}

func OfPtr[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Ptr
}

func OfSlice[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Slice
}

func OfStruct[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Struct
}

func OfUnsafePointer[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.UnsafePointer
}

func OfComplex64[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Complex64
}

func OfComplex128[A any](actual *A) bool {
	return reflect.TypeOf(*actual).Kind() == reflect.Complex128
}
