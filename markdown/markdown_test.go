package markdown

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cishiv/markdown-to-json/v2/utils"
)

func TestEndToEndParser(t *testing.T) {
	dat, err := os.ReadFile("./examples/small-test.md")
	if err != nil {
		panic(err)
	}
	result := fileToSlice(string(dat))
	patterns := compilePatterns()
	matchMap := match(patterns, result)
	postprocessedlines := precompute(matchMap)
	serializedLines := out(postprocessedlines)
	jsonString := utils.MapToJsonString(serializedLines)

	expected, err := os.Open("./examples/expected-result.json")

	if err != nil {
		panic(err)
	}

	byteResult, _ := ioutil.ReadAll(expected)
	var res map[int]Line
	json.Unmarshal([]byte(byteResult), &res)

	eq := true

	// TODO: Needs a better assertion.
	if len(res[0].Results) != len(serializedLines[0].Results) {
		eq = false
	}

	if !eq {
		t.Error("Result was not as expected")
		_ = ioutil.WriteFile("examples/test-result-fail.json", []byte(jsonString), 0644)
	} else {
		_ = ioutil.WriteFile("examples/test-result.json", []byte(jsonString), 0644)
	}
}
