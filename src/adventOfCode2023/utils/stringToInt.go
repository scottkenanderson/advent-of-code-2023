package utils

import "strconv"

func StringToInt(input string) int {
	gameIDStr := input
	gameID, _ := strconv.ParseInt(gameIDStr, 10, 0)
	return int(gameID)
}
