[![GoDoc](http://godoc.org/github.com/tdewolff/css?status.svg)](http://godoc.org/github.com/tdewolff/css)

~90% test coverage

# CSS

This package is a CSS3 tokenizer written in [Go][1] and follows the specification at [CSS Syntax Module Level 3](http://www.w3.org/TR/css3-syntax/). It takes an io.Reader and converts it into tokens until the EOF.

## Installation

Run the following command

	go get github.com/tdewolff/css

or add the following import and run project with `go get`

	import "github.com/tdewolff/css"

## Usage
The following initializes a new tokenizer with io.Reader `r`:

	z := css.NewTokenizer(r)

To tokenize until EOF an error, use:
``` go
for {
	tt := z.Next()
	switch tt {
	case css.ErrorToken:
		// error or EOF
		return
	// ...
	}
}
```

All tokens (see [CSS Syntax Module Level 3](http://www.w3.org/TR/css3-syntax/)):
``` go
ErrorToken
IdentToken
FunctionToken
AtKeywordToken
HashToken
StringToken
BadStringToken
UrlToken
BadUrlToken
DelimToken
NumberToken
PercentageToken
DimensionToken
UnicodeRangeToken
IncludeMatchToken
DashMatchToken
PrefixMatchToken
SuffixMatchToken
SubstringMatchToken
ColumnToken
WhitespaceToken
CDOToken
CDCToken
ColonToken
SemicolonToken
CommaToken
BracketToken // all bracket tokens use this, with Data() one can get the actual bracket
CommentToken
```

### Examples
Basic example:
``` go
package main

import (
	"os"

	"github.com/tdewolff/css"
)

// Tokenize CSS3 from stdin.
func main() {
	z := css.NewTokenizer(os.Stdin)
	for {
		tt := z.Next()
		switch tt {
		case css.ErrorToken:
			if z.Err() != io.EOF {
				fmt.Println("Error on line", z.Line(), ":", z.Err())
			}
			return
		case css.IdentToken:
			fmt.Println("Identifier", z.Data())
		case css.NumberToken:
			fmt.Println("Number", z.Data())
		// ...
		}
	}
}
```

[1]: http://golang.org/ "Go Language"
