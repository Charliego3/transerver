package utils

import (
	"database/sql"
	"github.com/gookit/goutil/strutil"
	"unsafe"
)

func String(data []byte) string {
	//hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	//return *(*string)(unsafe.Pointer(&reflect.StringHeader{
	//	Data: hdr.Data,
	//	Len:  hdr.Len,
	//}))
	return unsafe.String(&data[0], len(data))
}

func Bytes(data string) []byte {
	//hdr := *(*reflect.StringHeader)(unsafe.Pointer(&data))
	//return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
	//	Data: hdr.Data,
	//	Len:  hdr.Len,
	//	Cap:  hdr.Len,
	//}))
	return unsafe.Slice(unsafe.StringData(data), len(data))
}

// AnyBlank 存在任何一个字符串为空时返回true, 否则返回false
func AnyBlank(args ...string) bool {
	for _, v := range args {
		if strutil.IsBlank(v) {
			return true
		}
	}
	return false
}

// NonBlanks 有任何一个字符串为空时返回false, 否则返回true
func NonBlanks(args ...string) bool {
	for _, v := range args {
		if strutil.IsBlank(v) {
			return false
		}
	}
	return true
}

// Blanks 每一个值都为空时返回true, 存在任何一个不为空的字符串都返回false
func Blanks(args ...string) bool {
	for _, v := range args {
		if strutil.IsNotBlank(v) {
			return false
		}
	}
	return true
}

// SQLString returns a valid sql.NullString
func SQLString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func IsEmail(email string) bool {
	return false
}
