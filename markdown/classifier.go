package markdown

import "markdown-to-notion/utils"

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
				if nextResults[0] == NEWLINE {
					next = nextResults[0]
				}
			} else {
				if idx == 0 {
					// we might have spans IN blocks, just how we can have "block-esque" things in spans (this will likely require an index check)
					// if idx ALL spans > idx ALL blocks then BLOCK else SPAN
					if isNextLineBlock(spankeys, blockkeys, line.resultStrings, nextResults) {
						// this line is a block, it's likely the next line is going to be a paragraph_start if it not a new line or block, so check if its a block
						next = BLOCK_PATTERN
					} else {
						next = SPAN_PATTERN
					}
				}
			}
		}

		if (prev != SPAN_PATTERN && prev != NEWLINE) && computedLines[idx].safe {
			// we're in a paragraph
			computedLines[idx] = copyAndAddLineTypeToLinkedLine(computedLines[idx], "paragraph_internal")
		} else {
			if prev == NEWLINE && computedLines[idx].safe {
				// we're starting a paragraph
				computedLines[idx] = copyAndAddLineTypeToLinkedLine(computedLines[idx], "paragraph_start")
			} else {
				// this is something else, no paragraph
				if next == NEWLINE || next == BLOCK_PATTERN {
					computedLines[idx] = copyAndAddLineTypeToLinkedLine(computedLines[idx], "paragraph_end")
				} else {
					computedLines[idx] = copyAndAddLineTypeToLinkedLine(computedLines[idx], "block_start_end")
				}
			}
		}

		// change prev
		if len(line.resultStrings) == 1 {
			if line.resultStrings[0] == NEWLINE {
				prev = line.resultStrings[0]
			}
			// if the current line is not a span in any way, but is a block, then it is a block
			// otherwise, it must be a span (?)
		} else if !utils.ContainsAny(spankeys, line.resultStrings) &&
			utils.ContainsAny(blockkeys, line.resultStrings) {
			prev = BLOCK_PATTERN
		} else {
			prev = SPAN_PATTERN
		}

	}
	return computedLines
}

func isNextLineBlock(spankeys []string, blockkeys []string, resultStrings []string, nextResults []string) bool {
	return !utils.ContainsAny(spankeys, resultStrings) &&
		utils.ContainsAny(blockkeys, resultStrings) &&
		utils.ContainsAny(blockkeys, nextResults) &&
		!utils.ContainsAny(spankeys, nextResults)
}

func classifyLine() Classification {
	return PARAGRAPH_START
}
