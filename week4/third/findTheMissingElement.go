package third

//SortTheIntSlice is a function that sorts from minimum to maximum a slice of integers
func sortTheIntSlice(slice []int) (sortedslice []int) {
	for i := len(slice); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}

	}
	return slice
}

//FindMissingElement is a function that returns the single element that does not exist in both slices.
func FindMissingElement(firstSlice []int, secondSlice []int) (missingElement int) {
	sortedFirst := sortTheIntSlice(firstSlice)
	sortedSecond := sortTheIntSlice(secondSlice)
	var element int
	var check bool
	if len(sortedFirst) > len(sortedSecond) {
		for i := 0; i < len(sortedFirst)-1; i++ {
			if sortedFirst[i] != sortedSecond[i] {
				element = sortedFirst[i]
				i = len(sortedFirst)
				check = true
			}
			if check != true {
				element = sortedFirst[i+1]
			}
		}
	} else {
		for i := 0; i < len(sortedFirst); i++ {
			if sortedFirst[i] != sortedSecond[i] {
				element = sortedFirst[i]
				i = len(sortedFirst)
				check = true
			}
			if check != true {
				element = sortedSecond[i+1]
			}
		}
	}
	return element
}
