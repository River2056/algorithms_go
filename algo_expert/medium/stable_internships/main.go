package stableinternships

import "fmt"

func StableInternships(interns [][]int, teams [][]int) [][]int {
	chosenIntern := make(map[int]int)

	freeInterns := make([]int, len(interns))
	for i := range interns {
		freeInterns = append(freeInterns, i)
	}

	internInitialChoice := make([]int, len(interns))
	for range interns {
		internInitialChoice = append(internInitialChoice, 0)
	}

	teamsMap := make([]interface{}, 0)
	for _, teamRank := range teams {
		rank := make(map[int]int)
		for intern, ranking := range teamRank {
			rank[intern] = ranking
		}
		teamsMap = append(teamsMap, rank)
	}

	for len(freeInterns) > 0 {
		currentIntern := freeInterns[len(freeInterns)-1]
		freeInterns = freeInterns[:len(freeInterns)-1]
		currentInternChoice := internInitialChoice[currentIntern]
		internPreferredTeam := interns[currentIntern][currentInternChoice]

		if _, ok := chosenIntern[currentIntern]; !ok {
			chosenIntern[internPreferredTeam] = currentIntern
		} else {

		}
	}

	return [][]int{}
}
