// Parse and convert a trivialised markdown spec to an opinionated JSON format
package markdown

import (
	"fmt"
	"obsidian-to-notion/utils"
	"regexp"
	"strings"
)

type Pair[T, U any] struct {
	T any
	U any
}

type Match struct {
	name      string
	line      string
	lineIndex int
	indices   [][]int
}

type LinkedLine struct {
	lineType        string
	unparsedResults []Match
	resultStrings   []string
	safe            bool
	content         string
}

type Result struct {
	Matcher    string `json:"matcher"`
	IndexStart int    `json:"indexStart"`
	IndexEnd   int    `json:"indexEnd"`
}

type Line struct {
	Content string   `json:"content"`
	Results []Result `json:"results"`
	Type    string   `json:"type"`
}

// paragraph is fallback
var blockpatterns = map[string]string{
	// These matchers _should_ work because we're matching blocks line-by-line
	// recurrence capture doesn't bubble up the way I expect, this needs to be handled programatically
	// if something is an ###, it can't be ## or #
	"heading1":   "^#{1}",
	"heading2":   "^#{2}",
	"heading3":   "^#{3}",
	"linebreak":  "\n",
	"blockquote": "^>",          // support only single level blocks (blocks can contain other elements).
	"codeblock":  "^```",        // this one needs depth.
	"ul":         `^\-|^\*|^\+`, // this one is debatable.
	"ol":         `^[0-9]*\.`,
	"hr":         "^---",
}

// excludes images.
var spanpatterns = map[string]string{
	// for em and strong, we need to count, but regex works.
	"em":         "_.*_",
	"link":       `\[.*\]\(.*\)`,
	"strong":     `\*.*\*`,
	"inlinecode": "`.*`",
	"img":        "TODO", // TODO
}

// do not use - just for context / remembering that these might be useful things.
var miscpatterns = map[string]string{
	"escape":   "TODO", // TODO
	"autolink": "TODO", // TODO
}

// func Parse(markdownString string) Block {
// 	return Block{
// 		name:     "body",
// 		content:  "",
// 		children: [
// 			Block{},
// 			Block{}
// 		],
// 	}
// }

func fileToSlice(file string) []string {
	lines := strings.Split(file, "\n")
	return lines
}

func parse(fileSlice []string) string {
	// Parse line by line
	maxPtr := len(fileSlice) - 1
	for ptr := 0; ptr <= maxPtr; ptr++ {
		if fileSlice[ptr] == "" {
			continue
		}
		// fmt.Println(fileSlice[ptr])
	}
	return ""
}

func apply(patterns map[string]regexp.Regexp, line string, idx int) []Match {
	// fmt.Println(line)
	var matches []Match
	for name, pattern := range patterns {
		matched := pattern.MatchString(line)
		indices := pattern.FindAllIndex([]byte(line), -1)
		if matched {
			matches = append(matches, Match{
				name:      name,
				line:      line,
				indices:   indices,
				lineIndex: idx,
			})
		}
	}

	// append a no-op paragraph match here
	if len(matches) == 0 {
		matches = append(matches, Match{
			name:      "paragraph",
			line:      line,
			lineIndex: idx,
		})
	}
	return matches
}

func compilePatterns() map[string]regexp.Regexp {
	fmt.Println("compile block patterns")
	patterns := make(map[string]regexp.Regexp)
	for name, pattern := range blockpatterns {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		patterns[name] = *compiled.Copy()
		fmt.Println("Complied Name:", name, "=>", "Pattern:", pattern)
	}

	fmt.Println("compile span patterns")
	for name, pattern := range spanpatterns {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		patterns[name] = *compiled.Copy()
		fmt.Println("Compiled Name:", name, "=>", "Pattern:", pattern)
	}
	return patterns
}

// TODO: Refactor once working
func precompute(matchMap map[int][]Match) map[int]LinkedLine {
	// TODO: Key arrays should be constant
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

	// create computedLines map and preallocate memory for each line
	computedLines := make(map[int]LinkedLine)
	// precompute the things needed to determine hierarchies
	for index, results := range matchMap {
		// fmt.Println(line, results)
		resultStrings := utils.Map(results, func(t Match) string { return t.name })
		// if the previous thing is not a newline and the current thing doesn't contain lone block patterns, this thing is likely in a paragraph
		safe := true
		for _, m := range resultStrings {
			// does the current thing contain a block pattern without a span pattern?
			if !utils.Contains(spankeys, m) &&
				utils.Contains(blockkeys, m) {
				safe = !safe
				break
			}
		}
		// we have to ensure that the index used for the line is the one associated with the line
		// and not the internal ordering of the Matches map
		// it might be better to fix this by ensuring that we can encode the index in the key for matchMap
		// a traditional 'map' may not work here
		computedLines[index] = LinkedLine{
			resultStrings:   resultStrings,
			unparsedResults: results,
			safe:            safe,
			// We can reliably do this for the moment
			content: results[0].line,
		}
	}

	// for the lines computed, we can look back and forth based on the indices
	prev := "start"
	for idx, line := range computedLines {
		// we don't need to _store_ previous and next for a line, we only need it to determine if the line is
		// paragrah_start
		// paragraph_end
		// paragraph_internal
		// block
		next := ""
		if idx != len(computedLines)-1 {
			nextResults := computedLines[idx+1].resultStrings
			if len(nextResults) == 1 {
				if nextResults[0] == "newline" {
					next = nextResults[0]
				}
			} else {
				if idx == 0 {
					// we might have spans IN blocks, just how we can have "block-esque" things in spans (this will likely require an index check)
					// if idx ALL spans > idx ALL blocks then BLOCK else SPAN
					if !utils.ContainsAny(spankeys, line.resultStrings) &&
						utils.ContainsAny(blockkeys, line.resultStrings) &&
						utils.ContainsAny(blockkeys, nextResults) &&
						!utils.ContainsAny(spankeys, nextResults) {
						// this line is a block, it's likely the next line is going to be a paragraph_start if it not a new line or block, so check if its a block
						next = "block"
					} else {
						next = "span"
					}
				}
			}
		}

		if (prev != "block" && prev != "newline") && computedLines[idx].safe {
			// we're in a paragraph
			copied := LinkedLine{
				resultStrings:   computedLines[idx].resultStrings,
				unparsedResults: computedLines[idx].unparsedResults,
				safe:            computedLines[idx].safe,
				lineType:        "paragraph_internal",
				content:         computedLines[idx].content,
			}
			computedLines[idx] = copied
		} else {
			if prev == "newline" && computedLines[idx].safe {
				// we're starting a paragraph
				copied := LinkedLine{
					resultStrings:   computedLines[idx].resultStrings,
					unparsedResults: computedLines[idx].unparsedResults,
					safe:            computedLines[idx].safe,
					lineType:        "paragraph_start",
					content:         computedLines[idx].content,
				}
				computedLines[idx] = copied
			} else {
				// this is something else, no paragraph
				if next == "newline" || next == "block" {
					copied := LinkedLine{
						resultStrings:   computedLines[idx].resultStrings,
						unparsedResults: computedLines[idx].unparsedResults,
						safe:            computedLines[idx].safe,
						lineType:        "paragraph_end",
						content:         computedLines[idx].content,
					}
					computedLines[idx] = copied
				} else {
					copied := LinkedLine{
						resultStrings:   computedLines[idx].resultStrings,
						unparsedResults: computedLines[idx].unparsedResults,
						safe:            computedLines[idx].safe,
						lineType:        "block_start_end",
						content:         computedLines[idx].content,
					}
					computedLines[idx] = copied
				}
			}
		}

		// change prev
		if len(line.resultStrings) == 1 {
			if line.resultStrings[0] == "newline" {
				prev = line.resultStrings[0]
			}
			// if the current line is not a span in any way, but is a block, then it is a block
			// otherwise, it must be a span (?)
		} else if !utils.ContainsAny(spankeys, line.resultStrings) &&
			utils.ContainsAny(blockkeys, line.resultStrings) {
			prev = "block"
		} else {
			prev = "span"
		}

	}
	return computedLines
}
