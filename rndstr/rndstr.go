package rndstr

import (
	"math/rand"
	"regexp"
	"strings"
)

const (
	lower             = `abcdefghijklmnopqrstuvwxyz`
	upper             = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	numbers           = `0123456789`
	specials          = `!@#$%^*()_-+={}[]|:;,.?/`
	MaxPasswordLength = 64
	MinPasswordLength = 12
)

func IsComplex(s string) bool {
	for idx, c := range specials {
		if strings.Contains(s, string(c)) {
			break
		}

		if idx == len(s) {
			return false
		}
	}

	rules := []string{".{12,}", "[a-z]", "[A-Z]", "[0-9]"}
	for _, rule := range rules {
		r := regexp.MustCompile(rule)
		if !r.MatchString(s) {
			return false
		}
	}

	return true
}

func RandString(n int) string {
	if n > MaxPasswordLength {
		n = MinPasswordLength
	}

	sb := strings.Builder{}
	all := []string{lower, upper, numbers, specials}

	for i, g := 0, 0; i < n; i, g = i+1, rand.Intn(4) {
		x := rand.Intn(len(all[g]))
		sb.WriteByte(all[g][x])
	}

	return sb.String()
}
