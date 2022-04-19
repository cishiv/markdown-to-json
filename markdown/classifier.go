package markdown

import (
	"strings"

	"github.com/cishiv/markdown-to-json/v2/utils"
)

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
		linkedLine := LinkedLine{
			resultStrings:   resultStrings,
			unparsedResults: results,
			safe:            safe,
			// We can reliably do this for the moment
			content: results[0].line,
		}
		lineType := isBlockOrNewline(linkedLine, blockkeys, spankeys)
		copied := copyAndAddLineTypeToLinkedLine(linkedLine, string(lineType))
		computedLines[index] = copied
	}

	for idx, line := range computedLines {
		if line.lineType == string(CONTEXTUAL_CLASSIFICATION) {
			if idx == len(computedLines)-1 {
				computedLines[idx] = copyAndAddLineTypeToLinkedLine(line, string(PARAGRAPH_END))
			} else if idx == 0 {
				computedLines[idx] = copyAndAddLineTypeToLinkedLine(line, string(PARAGRAPH_START))
			} else {
				classification := classifyLine(line, blockkeys, spankeys, computedLines[idx-1].lineType, computedLines[idx+1].lineType)
				computedLines[idx] = copyAndAddLineTypeToLinkedLine(line, string(classification))
			}
		}
	}

	// cull extranous heading matches
	m := map[string]int{
		"heading3": 0,
		"heading2": 1,
		"heading1": 2,
	}

	for idx, line := range computedLines {
		if line.lineType == string(BLOCK) {
			foundResult := Match{}
			lowest := int(^uint(0) >> 1)
			for _, result := range line.unparsedResults {
				if result.name == "heading3" || result.name == "heading2" || result.name == "heading1" {
					if m[result.name] < lowest {
						foundResult = result
					}
				}
			}
			var newResults []Match
			newResults = append(newResults, foundResult)
			computedLines[idx] = LinkedLine{
				lineType:        line.lineType,
				unparsedResults: newResults,
				resultStrings:   line.resultStrings,
				safe:            line.safe,
				content:         line.content,
			}
		}
	}
	return computedLines
}

// context free classifications
func isBlockOrNewline(line LinkedLine, blockkeys []string, spankeys []string) Classification {
	// base case
	// this might be breaking
	if strings.Contains(line.content, "idx:") {
		return NEWLINE
	}
	// block logic - needs a unit test
	if utils.ContainsAny(blockkeys, line.resultStrings) && !utils.ContainsAny(spankeys, line.resultStrings) {
		return BLOCK
		// if we hit both, we need to check that the block occurs first
	} else if utils.ContainsAny(blockkeys, line.resultStrings) && utils.ContainsAny(spankeys, line.resultStrings) {
		globalLowest, globalBlock := int(^uint(0)>>1), false
		for _, match := range line.unparsedResults {
			// first check if the name of the pattern is a block
			localLowest, block := int(^uint(0)>>1), utils.Contains(blockkeys, match.name)
			// get the lowest found index in the indices of this matcher
			for _, indexpair := range match.indices {
				if indexpair[0] < localLowest {
					localLowest = indexpair[0]
				}
			}
			// bubble up the local maxima and whether it was for a block or not
			if localLowest < globalLowest {
				globalLowest, globalBlock = localLowest, block
			}
		}
		if globalBlock {
			return BLOCK
		}
	}
	return CONTEXTUAL_CLASSIFICATION
}

// contextual classification
func classifyLine(line LinkedLine, blockkeys []string, spankeys []string, previous string, next string) Classification {
	if previous == string(BLOCK) || previous == string(NEWLINE) {
		return PARAGRAPH_START
	} else if (previous == string(PARAGRAPH_INTERNAL) || previous == string(PARAGRAPH_START)) &&
		(next != string(BLOCK) && next != string(NEWLINE)) {
		return PARAGRAPH_INTERNAL
	} else if previous == string(CONTEXTUAL_CLASSIFICATION) || next == string(CONTEXTUAL_CLASSIFICATION) {
		return PARAGRAPH_INTERNAL
	} else {
		return PARAGRAPH_END
	}
}
