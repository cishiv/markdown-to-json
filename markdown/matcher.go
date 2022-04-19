package markdown

import (
	"regexp"
	"strconv"

	"github.com/cishiv/markdown-to-json/v2/utils"
)

func apply(patterns map[string]regexp.Regexp, line string, idx int) []Match {
	var matches []Match
	for name, pattern := range patterns {
		// trim leading and trailing space so we can match against indented lists, we maintain the space in the actual line "content" val
		trimmed, offset := utils.TrimAndCount(line)
		matched := pattern.MatchString(trimmed)
		paddedIndices := pattern.FindAllIndex([]byte(trimmed), -1)
		// reapply the offset to each index pair
		var indices = utils.Matrix2D[int](len(paddedIndices), 2)
		for i, pair := range paddedIndices {
			indices[i][0] = pair[0] + offset
			indices[i][1] = pair[1] + offset
		}
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
			name:      PARAGRAPH,
			line:      line,
			lineIndex: idx,
		})
	}
	return matches
}

func compilePatterns() map[string]regexp.Regexp {
	patterns := make(map[string]regexp.Regexp)
	for name, pattern := range blockpatterns {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		patterns[name] = *compiled.Copy()
	}

	for name, pattern := range spanpatterns {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		patterns[name] = *compiled.Copy()
	}
	return patterns
}

func match(patterns map[string]regexp.Regexp, lines []string) map[int][]Match {
	matchMap := make(map[int][]Match)
	for index, line := range lines {
		if line == "" {
			// unique keys for newlines
			var matches []Match
			matches = append(matches, Match{
				name:      string(NEWLINE),
				line:      line + "idx:" + strconv.Itoa(index),
				lineIndex: index,
			})
			matchMap[index] = matches
		} else {
			matches := apply(patterns, line, index)
			matchMap[index] = matches
		}
	}
	return matchMap
}
