package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	var item int
	var result bool
	if index > cap(slice)-1 || index < 0 {
		item = 0
		result = false
	} else {
		item = slice[index]
		result = true
	}
	return item, result
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	var newSlice []int = make([]int, cap(slice))
	copy(newSlice, slice)
	if index > cap(newSlice)-1 || index < 0 {
		newSlice = append(newSlice, value)
	} else {
		newSlice[index] = value
	}
	return newSlice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var slice []int
	if length > 0 {
		slice = make([]int, length)
		for i := 0; i < length; i++ {
			slice[i] = value
		}
	} else {
		slice = make([]int, 0)
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var newSlice []int = make([]int, cap(slice))
	copy(newSlice, slice)
	if index > cap(newSlice)-1 || index < 0 {
		return newSlice
	} else {
		newSlice = append(newSlice[:index], newSlice[index+1:]...)
		return newSlice
	}
}
