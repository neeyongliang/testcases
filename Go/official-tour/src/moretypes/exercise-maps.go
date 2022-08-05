package moretypes

import (
	"golang-tour/mywc"
	"strings"
)

func WordCount(s string) map[string]int {
	stringsplits := strings.Fields(s)
	m := make(map[string]int)
	for _, value := range stringsplits {
		m[value] += 1
	}
	return m
}

func ExerciseMaps() {
	mywc.Test(WordCount)
}