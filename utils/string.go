package utils

import (
	"github.com/gookit/goutil/strutil"
	"reflect"
	"unsafe"
)

func String(data []byte) string {
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
	}))
}

func Bytes(data string) []byte {
	hdr := *(*reflect.StringHeader)(unsafe.Pointer(&data))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
		Cap:  hdr.Len,
	}))
}

func NonBlanks(v ...string) bool {
	for _, i := range v {
		if strutil.IsBlank(i) {
			return false
		}
	}
	return true
}

func Blanks(v ...string) bool {
	for _, i := range v {
		if strutil.IsNotBlank(i) {
			return false
		}
	}
	return true
}
