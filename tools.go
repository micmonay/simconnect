package simconnect

import (
	"errors"
	"unsafe"
)

func convStrToGoString(buf []byte) string {
	var index int
	var value byte
	for index, value = range buf {
		if value == 0x00 {
			break
		}
	}
	return string(buf[:index])

}

func convGoStringtoBytes(str string) []byte {
	str = str + "\x00"
	return []byte(str)
}

func convCBytesToGoBytes(ptr unsafe.Pointer, size int) ([]byte, error) {
	if size > 1<<30 {
		return nil, errors.New("Dispatch return to big size array data")
	}
	buf := make([]byte, size)
	copy(buf, (*[1 << 30]byte)(ptr)[:size:size])
	return buf, nil
}
