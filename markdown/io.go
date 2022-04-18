package markdown

import (
	"markdown-to-json/utils"
	"strings"
)

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

func out(linkedLines map[int]LinkedLine) map[int]Line {
	lines := make(map[int]Line)
	for idx, v := range linkedLines {
		results := utils.Map(v.unparsedResults, func(t Match) Result {
			var occurences []Ocurrence
			for _, val := range t.indices {
				occurences = append(occurences, Ocurrence{
					FirstIdx:  val[0],
					SecondIdx: val[1],
				})
			}
			return Result{
				Matcher:    t.name,
				Occurences: occurences,
			}
		})
		lines[idx] = Line{
			Content: v.content,
			Results: results,
			Type:    v.lineType,
		}
	}
	return lines
}
