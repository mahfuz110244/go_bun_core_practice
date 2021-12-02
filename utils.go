package main

import (
	"strconv"
)

func getInt(data string) int {
	i, err := strconv.Atoi(data)
	if err != nil {
		return 0
	}
	return i
}
