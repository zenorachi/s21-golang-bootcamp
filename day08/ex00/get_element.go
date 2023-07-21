package ptr

import (
	"errors"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if !isValidData(arr, idx) {
		return 0, errors.New("invalid input data")
	}

	ptr := unsafe.Pointer(&arr[0])
	step := unsafe.Sizeof(arr[0])
	ptr = unsafe.Pointer(uintptr(ptr) + step*uintptr(idx))

	result := *(*int)(ptr)
	return result, nil
}

func isValidData(arr []int, idx int) bool {
	length := len(arr)
	if idx >= length || idx < 0 || length == 0 || arr == nil {
		return false
	}
	return true
}
