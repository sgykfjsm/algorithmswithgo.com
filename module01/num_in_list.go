package module01

import (
	"sort"
)

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if len(list) == 0 {
		return false
	}

	sort.Ints(list)

	for {
		idx := len(list) / 2

		if list[idx] == num {
			return true
		}

		if len(list) == 1 {
			break
		}

		if list[idx] > num {
			list = list[:len(list)/2]
		} else {
			list = list[len(list)/2:]
		}
	}

	return false
}
