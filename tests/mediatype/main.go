// +build gofuzz
package fuzz

import "github.com/tdewolff/parse"

func Fuzz(data []byte) int {
	_, _ = parse.Mediatype(data)
	return 1
}
