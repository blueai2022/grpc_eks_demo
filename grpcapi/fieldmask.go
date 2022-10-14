package grpcapi

import "strings"

// A function that maps field mask field names to the names used in Go structs.
// It has to be implemented according to your needs.
func naming(s string) string {
	return strings.Title(s)
}
