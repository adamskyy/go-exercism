package atbashcipher

import "strings"

func Atbash(s string) string {
	output := []rune{}
    s = strings.ToLower(s)
    written := 0
	for _, v := range s {
        if !(v >= '0' && v <= '9') && !(v >= 'a' && v <= 'z') {
            continue
        }
        if written > 0 && written%5 == 0 {
            output = append(output, ' ')
        }
        if v >= '0' && v <= '9' {
            output = append(output, v)
        } else if v >= 'a' && v <= 'z' {
            t := v - 'a'
            output = append(output, 'z' - t)
        }
        written++
	}
	return string(output)
}