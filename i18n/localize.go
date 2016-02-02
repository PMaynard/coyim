package i18n

import "fmt"

// T marks a string literal as transatable
type T string

// Local returns the given string in the local language
func Local(v string) string {
	return v
}

// Localf returns the given string in the local language. It supports Printf formatting.
func Localf(f string, p ...interface{}) string {
	return fmt.Sprintf(f, p...)
}
