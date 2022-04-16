package markdown

import (
	"os"
	"testing"
)

var unparsedMarkdown string = `
# This is a heading
## This is a second heading
- This is a list item
- This is also a list item
This should be a paragraph
---
**bold**
_italics_
[Link](https://link.com)
`

func TestFileToSlice(t *testing.T) {
	dat, err := os.ReadFile("./test.md")
	if err != nil {
		panic(err)
	}
	result := fileToSlice(string(dat))
	got := len(result)
	parse(result)
	want := 310
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPatterns(t *testing.T) {
	dat, err := os.ReadFile("./test.md")
	if err != nil {
		panic(err)
	}
	result := fileToSlice(string(dat))
	patterns := compilePatterns()
	// o(n^2) nice :)
	for _, line := range result {
		apply(patterns, line)
	}
}
