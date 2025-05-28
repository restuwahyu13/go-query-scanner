package helper

import "regexp"

func CheckValidQuery(query string) bool {
	pattern := `^(([^#&?]*)=([^#&]*))?(&([^#&?]*)=([^#&]*))*[#]?$`
	regex := regexp.MustCompile(pattern)

	if ok := regex.MatchString(query); !ok {
		return false
	}

	return true
}
