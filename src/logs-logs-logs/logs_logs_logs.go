package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, s := range log {
		if s == 10071 {
			return "recommendation"
		} else if s == 128269 {
			return "search"
		} else if s == 9728 {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""
	for _, s := range log {
		if s == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(s)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	for i, s := range log {
		if s == 10071 && i < limit {
			return true
		} else if s == 128269 && i < limit {
			return true
		} else if s == 9728 && i < limit {
			return true
		}
	}
	return false

}
