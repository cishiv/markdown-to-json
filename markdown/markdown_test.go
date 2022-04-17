package markdown

import (
	"fmt"
	"obsidian-to-notion/utils"
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
	matchMap := make(map[string][]Match)
	for _, line := range result {
		if line == "" {
			var matches []Match
			matches = append(matches, Match{
				name: "newline",
				line: line,
			})
			matchMap[line] = matches
		} else {
			matches := apply(patterns, line)
			matchMap[line] = matches
		}
	}

	prev := "start"
	// blockpattern names
	// pre-allocation because len(blockpatterns) is not going to change
	spankeys := make([]string, len(spanpatterns))
	i := 0
	for k := range spanpatterns {
		spankeys[i] = k
		i++
	}

	blockkeys := make([]string, len(blockpatterns))
	j := 0
	for k := range blockpatterns {
		blockkeys[j] = k
		j++
	}

	for line, results := range matchMap {
		resultStrings := utils.Map(results, func(t Match) string { return t.name })
		// if the previous thing is not a newline and the current thing doesn't contain lone block patterns, this thing is likely in a paragraph
		safe := true
		for _, m := range resultStrings {
			// does the current thing contain a block pattern without a span pattern?
			if !utils.Contains(spankeys, m) && utils.Contains(blockkeys, m) {
				safe = !safe
			}
		}
		if prev != "newline" && safe {
			// we're in a paragraph
			fmt.Println("in a paragraph")
		} else {
			if prev == "newline" && safe {
				// we're starting a paragraph
				fmt.Println("starting a new paragraph")
			} else {
				// this is something else, no paragraph
				fmt.Println("likely something else")
			}
		}
		fmt.Println(line, results)
	}
}
