package grpcapi

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	titleCaser = cases.Title(language.Und, cases.NoLower)
)

// A function that maps field mask field names to the names used in Go structs.
// It has to be implemented according to your needs.
func naming(s string) string {
	return titleCaser.String(s) //strings.Title(s)
}
