// +build gofuzz
package fuzz

import "github.com/tdewolff/parse"

func Fuzz(data []byte) int {
	_, _, _ = parse.DataURI(data)
	return 1
}
