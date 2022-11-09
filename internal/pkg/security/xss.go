package security

import (
	"github.com/microcosm-cc/bluemonday"
)

var Policy *bluemonday.Policy

func init() {
	Policy = bluemonday.UGCPolicy()
}

// Sanitize return sanitized string
func Sanitize(input string) string {
	return Policy.Sanitize(input)
}
