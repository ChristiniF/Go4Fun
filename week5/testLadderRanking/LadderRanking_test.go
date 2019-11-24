package test

import "testing"

func TestLadderRanking(t *testing.T) {
	var rankScores []int
	result := []int{6, 5, 4}
	myScores := []int{3, 25, 50}
	currentLeaders := []int{110, 110, 80, 60, 60, 30, 30, 25}
	rankScores = LadderRanking(currentLeaders, myScores)
	if result[0] != rankScores[0] && result[1] != rankScores[1] && result[2] != rankScores[2] {
		t.Errorf("Incorrect")
	}
}

func TestLadderRanking2(t *testing.T) {
	var rankScores []int
	result := []int{4, 2, 1}
	myScores := []int{50, 125, 550}
	currentLeaders := []int{110, 500, 80}
	rankScores = LadderRanking(currentLeaders, myScores)
	if result[0] != rankScores[0] && result[1] != rankScores[1] && result[2] != rankScores[2] {
		t.Errorf("Incorrect")
	}
}

//LadderRanking function returns a slice of integers that represent actual rank in the leaderboard after each game.
func LadderRanking(currentLeaders []int, myScores []int) []int {

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
