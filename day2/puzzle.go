package day2

// Part1 What would your total score be if everything goes exactly according to your strategy guide?
func Part1() (totalScore int) {
	for round := range getRounds() {
		totalScore += scoreRound(round)
	}
	return totalScore
}
