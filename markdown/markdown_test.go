package markdown

import (
	"fmt"
	"obsidian-to-notion/utils"
	"os"
	"strconv"
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

// func TestFileToSlice(t *testing.T) {
// 	dat, err := os.ReadFile("./test.md")
// 	if err != nil {
// 		panic(err)
// 	}
// 	result := fileToSlice(string(dat))
// 	got := len(result)
// 	parse(result)
// 	want := 310
// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

func TestPatterns(t *testing.T) {
	dat, err := os.ReadFile("./test.md")
	if err != nil {
		panic(err)
	}
	result := fileToSlice(string(dat))
	patterns := compilePatterns()
	// o(n^2) nice :)
	matchMap := make(map[int][]Match)
	for index, line := range result {
		if line == "" {
			// unique keys for newlines
			var matches []Match
			matches = append(matches, Match{
				name:      "newline",
				line:      line + "idx:" + strconv.Itoa(index),
				lineIndex: index,
			})
			matchMap[index] = matches
		} else {
			matches := apply(patterns, line, index)
			matchMap[index] = matches
		}
	}
	postprocessedlines := precompute(matchMap)
	fmt.Println(utils.MapToJsonString(postprocessedlines))
}
