package anagram

import (
	"strings"
	"unicode/utf8"
)

func Detect(subject string, candidates []string) []string {
	output := []string{}

	target := strings.ToLower(subject)
	targetCount := utf8.RuneCountInString(target)
	targetFreqMap := buildTargFreqMap(target)

	for _, candidate := range candidates {

		candidateLower := strings.ToLower(candidate)
		candidateCount := utf8.RuneCountInString(candidateLower)

		if candidateCount != targetCount {
			continue
		} else if candidateLower == target {
			continue
		}

		if compareCandidateWithTarget(targetFreqMap, candidateLower) {
			output = append(output, candidate)
		}

	}
	return output
}

func buildTargFreqMap(word string) map[rune]int {
	subjectMapped := map[rune]int{}
	for _, v := range word {
		subjectMapped[v]++
	}
	return subjectMapped
}

func compareCandidateWithTarget(targetFreqMap map[rune]int, candidate string) bool {
	candidateMapped := map[rune]int{}
	for _, v := range candidate {
		candidateMapped[v]++
		val, ex := targetFreqMap[v]
		if !ex {
			return false
		}
		if candidateMapped[v] > val {
			return false
		}
	}
	return true
}
