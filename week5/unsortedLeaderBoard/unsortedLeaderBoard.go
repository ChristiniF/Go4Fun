package main

import (
	"fmt"
)

func main() {
	var rankScores []int
	randomScores := []int{3, 25, 120, 50, 80, 90, 110, 120}
	randomLeaders := []int{110, 110, 80, 60, 60, 30, 25}
	rankScores = ladderRanking(randomLeaders, randomScores)
	fmt.Println("rank", rankScores)
}

func ladderRanking(leaders []int, scores []int) []int {

	var slice []int
	var rank []int
	currentLeaders := sortMaxMinSlice(leaders)
	myScores := sortMinMaxSlice(scores)
	for i := 0; i < (len(currentLeaders) - 1); i++ {
		if currentLeaders[i] == currentLeaders[i+1] {
			slice = append(slice, currentLeaders[i+1])
			i++
		} else {
			slice = append(slice, currentLeaders[i])
		}
	}

	if currentLeaders[len(currentLeaders)-2] != currentLeaders[len(currentLeaders)-1] {
		slice = append(slice, currentLeaders[len(currentLeaders)-1])
	}
	for i := 0; i < len(myScores); i++ {
		for j := 0; j < len(slice); j++ {
			if myScores[i] >= slice[j] {
				rank = append(rank, j+1)
				j = len(slice)
			}
			if j == (len(slice)-1) && myScores[i] < slice[j] {
				rank = append(rank, j+2)
			}
		}
	}
	return rank
}

//sortMinMaxSlice is a function that sorts from minimum to maximum a slice of integers
func sortMinMaxSlice(slice []int) (sortedslice []int) {
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

//sortMaxMinSlice is a function that sorts from maximum to minimum a slice of integers
func sortMaxMinSlice(slice []int) (sortedslice []int) {
	for i := len(slice); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if slice[j] < slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}

	}
	return slice
}
