module github.com/tdewolff/fuzz/minify/replace-entities

go 1.13

replace github.com/tdewolff/parse => ../../../parse

require (
	github.com/dvyukov/go-fuzz v0.0.0-20191022152526-8cb203812681 // indirect
	github.com/tdewolff/parse v2.3.10
)
