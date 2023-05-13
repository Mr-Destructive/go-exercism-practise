package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	match := re.FindString(text)
	if match != "" {
		return true
	} else {
		return false
	}
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<([~=*-]+)?>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`"(.*?(?i)password.*?)"`)
	count := 0
	for _, line := range lines {
		c := re.Match([]byte(line))
		if c {
			count++
		}

	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line([0-9])*`)
	return string(re.ReplaceAll([]byte(text), []byte("")))
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+(\w+) `)
	new_lines := lines
	for i, line := range lines {
		if re.Match([]byte(line)) {
			user := strings.Split(re.FindString(line), "User ")
			new_lines[i] = "[USR] " + strings.Trim(user[1], " ") + " " + line
		}
	}
	return new_lines
}
