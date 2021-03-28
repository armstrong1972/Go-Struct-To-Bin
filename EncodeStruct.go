package main

// "unsafe"
import (
	"reflect"
	"unsafe"
)

func EncodeStruct(obj interface{}) ([]byte, error) {

	type sliceMock struct {
		addr uintptr
		len  int
		cap  int
	}

	typ := reflect.TypeOf(obj)
	size := typ.Elem().Size()

	oStruct := reflect.ValueOf(obj).Elem()

	binHead := &sliceMock{
		addr: uintptr(unsafe.Pointer(oStruct.UnsafeAddr())),
		cap:  int(size),
		len:  int(size),
	}

	bin := *(*[]byte)(unsafe.Pointer(binHead))

	return bin, nil
}
