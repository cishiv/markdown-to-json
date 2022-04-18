package markdown

import (
	"fmt"
	"regexp"
	"strconv"
)

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
			name:      PARAGRAPH,
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

func match(patterns map[string]regexp.Regexp, lines []string) map[int][]Match {
	matchMap := make(map[int][]Match)
	for index, line := range lines {
		if line == "" {
			// unique keys for newlines
			var matches []Match
			matches = append(matches, Match{
				name:      NEWLINE,
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
