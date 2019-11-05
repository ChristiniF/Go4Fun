package second

//FindMissingElement is a function that returns the single element that does not exist in both slices.
func FindMissingElement(firstSlice []int, secondSlice []int) (missingElement int) {
	var element int
	var sumFirstSlice int
	var sumSecondSlice int
	for _, v := range firstSlice {
		sumFirstSlice += v
	}
	for _, w := range secondSlice {
		sumSecondSlice += w
	}
	if len(firstSlice) > len(secondSlice) {
		element = sumFirstSlice - sumSecondSlice
	} else {
		element = sumSecondSlice - sumFirstSlice
	}

	return element
}
