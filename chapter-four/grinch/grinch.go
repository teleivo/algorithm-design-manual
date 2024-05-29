// Package set solves exercise 4.9 4-1.
package grinch

import (
	"slices"
)

// partition splits players into two teams of which the first contains the best while the second
// contains the worst players. An additional advantage is given to the best team if the number of
// players is odd by giving them one extra player.
// Time: O(N log N)
// Space: O(N) - copying the players into extra slice to sort them
func partition(players []int) ([]int, []int) {
	sortedPlayers := make([]int, len(players))
	copy(sortedPlayers, players)
	slices.Sort(sortedPlayers)

	var worst, best []int
	if len(players)%2 != 0 {
		worst = make([]int, len(players)/2)
		best = make([]int, len(players)/2+1)
	} else {
		worst = make([]int, len(players)/2)
		best = make([]int, len(players)/2)
	}

	for i, p := range sortedPlayers {
		if i < len(worst) {
			worst[i] = p
		} else {
			best[i-len(worst)] = p
		}
	}

	return best, worst
}
