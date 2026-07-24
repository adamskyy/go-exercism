package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
    re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	separator := regexp.MustCompile(`<[~*=-]*>`)
	return separator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	quotedPassword := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)

	count := 0
	for _, line := range lines {
		if quotedPassword.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(line string) string {
	endOfLine := regexp.MustCompile(`end-of-line\d+`)
	return endOfLine.ReplaceAllString(line, "")

}

func TagWithUserName(lines []string) []string {
	userPattern := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		matches := userPattern.FindStringSubmatch(line)
		if len(matches) > 1 {
			username := matches[1]
			lines[i] = "[USR] " + username + " " + line
		}
	}
	return lines
}
