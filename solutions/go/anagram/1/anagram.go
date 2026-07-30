package anagram

import (
	"strings"
	"fmt"
)

func Detect(subject string, candidates []string) []string {
    output := []string{}
	subjectMapped := map[rune]int{}
    subject = strings.ToLower(subject)
    for _, v := range subject {
        subjectMapped[v]++
    }
    fmt.Println(subjectMapped)
    for _, candidate := range candidates {
        invalid := false
        // Make a copy of target
        tempTarget := map[rune]int{}
        for k, v := range subjectMapped {
            tempTarget[k] = v
        }
        candidateTemp := strings.ToLower(candidate)
        if subject == candidateTemp {
            continue
        }
        for _, char := range candidateTemp {
            // Check whether current rune was scanned in the target word, if not go to the next word
            if _, ok := tempTarget[char]; !ok {
                invalid = true
                break
            } else {
                tempTarget[char]--
                if tempTarget[char] < 0 {
                    invalid = true
                    break
                }
            }
        }
        fmt.Println(tempTarget)
        if invalid {
            continue
        }
        allZeros := true
        for _, v := range tempTarget {
            if v != 0 {
                allZeros = false
                break
            }
        }
        if allZeros {
            output = append(output, candidate)
        }   
    }
    return output
}
