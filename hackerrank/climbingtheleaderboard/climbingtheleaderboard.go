// Package climbingtheleaderboard provides a solution implementation for the Climbing the Leaderboard problem.
//
// Statement: An arcade game player wants to climb to the top of the leaderboard and track their ranking.
// Given the leaderboard scores and the player's scores, find the player's rank after each new score.
//
// A complete description of this problem can be found here: https://www.hackerrank.com/challenges/climbing-the-leaderboard
package climbingtheleaderboard

// ClimbingLeaderboard finds a player's rank after new scores.
func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {

	rankLen := len(ranked)
	playerScoreLen := len(player)
	result := make([]int32, playerScoreLen)

	idRank := 0
	playerPosScore := int32(1)
	rankedScore := ranked[idRank]

	// Iterating player scores from back to front.
	for idPlayerScore := playerScoreLen - 1; idPlayerScore >= 0; idPlayerScore-- {
		playerScore := player[idPlayerScore]

		for (playerScore < rankedScore) && (idRank < rankLen) {

			idRank++
			// Update player score position after all ranked score has been checked.
			if idRank == rankLen {
				playerPosScore++
				break
			}

			rankedScore = ranked[idRank]
			// Only update playerPosScore when new rankedScore value is different than previous value.
			if rankedScore < ranked[idRank-1] {
				playerPosScore++
			}
		}

		// Setting the rank position of the current player score.
		result[idPlayerScore] = playerPosScore
	}

	return result
}
