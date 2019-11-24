package main

import (
	"fmt"
)

func main() {
	var rankScores []int
	myScores := []int{3, 25, 50, 80, 90, 110, 120}
	currentLeaders := []int{110, 110, 80, 60, 60, 30, 30, 25}
	rankScores = ladderRanking(currentLeaders, myScores)
	fmt.Println("rank", rankScores)
}

func ladderRanking(currentLeaders []int, myScores []int) []int {

	var slice []int
	var rank []int

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
