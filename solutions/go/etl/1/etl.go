package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    output := make(map[string]int)
	for k, v := range in {
        for _, letter := range v {
            output[strings.ToLower(letter)] = k
        }
    }
    return output
}
