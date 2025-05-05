package secretmanager_paramstore

import (
	"regexp"
	"strings"
)

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func pascalToKebab(str string) string {
	snake := matchAllCap.ReplaceAllString(str, "${1}-${2}")
	return strings.ToLower(snake)
}
