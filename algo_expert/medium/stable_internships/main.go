package stableinternships

func StableInternships(interns [][]int, teams [][]int) [][]int {
chosenIntern := make(map[int]int)

	freeInterns := make([]int, 0)
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
		for ranking, intern := range teamRank {
			rank[intern] = ranking
		}
		teamsMap = append(teamsMap, rank)
	}

	for len(freeInterns) > 0 {
		currentIntern := freeInterns[len(freeInterns)-1]
		freeInterns = freeInterns[:len(freeInterns)-1]
		currentInternChoice := internInitialChoice[currentIntern]
		internInitialChoice[currentIntern]++
		internPreferredTeam := interns[currentIntern][currentInternChoice]

		if _, ok := chosenIntern[internPreferredTeam]; !ok {
			chosenIntern[internPreferredTeam] = currentIntern
		} else {
			assignedIntern := chosenIntern[internPreferredTeam]
            assignInternRank := teamsMap[internPreferredTeam].(map[int]int)[assignedIntern]
            currentInternRank := teamsMap[internPreferredTeam].(map[int]int)[currentIntern]
            if currentInternRank < assignInternRank {
                chosenIntern[internPreferredTeam] = currentIntern
                freeInterns = append(freeInterns, assignedIntern)
            } else {
                freeInterns = append(freeInterns, currentIntern)
            }
		}
	}

	result := make([][]int, 0)
	for key, value := range chosenIntern {
		result = append(result, []int{value, key})
	}

	return result
}
