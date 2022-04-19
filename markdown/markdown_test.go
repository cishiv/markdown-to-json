package markdown

import (
	"io/ioutil"
	"markdown-to-json/utils"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	dat, err := os.ReadFile("./examples/small-test.md")
	if err != nil {
		panic(err)
	}
	result := fileToSlice(string(dat))
	patterns := compilePatterns()
	// o(n^2) nice :)
	matchMap := match(patterns, result)
	postprocessedlines := precompute(matchMap)
	serializedLines := out(postprocessedlines)
	jsonString := utils.MapToJsonString(serializedLines)
	_ = ioutil.WriteFile("examples/test-result.json", []byte(jsonString), 0644)
}
