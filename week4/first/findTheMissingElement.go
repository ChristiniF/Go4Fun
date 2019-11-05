package first

//FindMissingElement is a function that returns the single element that does not exist in both slices.
func FindMissingElement(firstSlice []int, secondSlice []int) (missingElement int) {
	var element int
	var equalFound int
	if len(firstSlice) > len(secondSlice) {
		for i := 0; i < len(firstSlice); i++ {
			for j := 0; j < len(secondSlice); j++ {
				if firstSlice[i] == secondSlice[j] {
					equalFound = firstSlice[i]
					j = len(secondSlice)
				}
			}
			if firstSlice[i] != equalFound {
				element = firstSlice[i]
				i = len(firstSlice)

			}
		}
	} else {
		for i := 0; i < len(secondSlice); i++ {
			for j := 0; j < len(firstSlice); j++ {
				if secondSlice[i] == firstSlice[j] {
					equalFound = secondSlice[i]
					j = len(firstSlice)
				}
			}
			if secondSlice[i] != equalFound {
				element = secondSlice[i]
				i = len(secondSlice)

			}
		}

	}

	return element
}
