// Parse and convert a trivialised markdown spec to an opinionated JSON format
package markdown

import (
	"io/ioutil"
	"os"

	"github.com/cishiv/markdown-to-json/v2/utils"
)

type Ocurrence struct {
	FirstIdx  int `json:"firstIdx"`
	SecondIdx int `json:"secondIdx"`
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

type Line struct {
	Content string   `json:"content"`
	Results []Result `json:"results"`
	Type    string   `json:"type"`
}

type Result struct {
	Matcher    string      `json:"matcher"`
	Occurences []Ocurrence `json:"occurences"`
}

// From Markdown
func FromMarkdownFileToJsonString(markdownFilePath string) string {
	dat, err := os.ReadFile(markdownFilePath)
	if err != nil {
		panic(err)
	}
	return FromMarkdownStringToJsonString(string(dat))
}

func FromMarkdownStringToJsonString(markdownString string) string {
	return toMarkdown(markdownString)
}

func FromMarkdownFileToJsonFile(markdownFilePath string, jsonFilePath string) {
	jsonString := FromMarkdownFileToJsonString(markdownFilePath)
	_ = ioutil.WriteFile(jsonFilePath, []byte(jsonString), 0644)
}

func FromMarkdownStringToJsonFile(markdownString string, jsonFilePath string) {
	jsonString := toMarkdown(markdownString)
	_ = ioutil.WriteFile(jsonFilePath, []byte(jsonString), 0644)
}

// To Markdown -  TODO
func ToMarkdownStringFromJsonString(jsonString string) string {
	return ""
}

func ToMarkdownFileFromJsonString(jsonString string) string {
	return ""
}

func ToMarkdownStringFromJsonFile(markdownFilePath string, jsonFilePath string) {
}

func ToMarkdownFileFromJsonFile(markdownFilePath string, jsonFilePath string) {
}

func toMarkdown(markdownString string) string {
	result := fileToSlice(markdownString)
	patterns := compilePatterns()
	matchMap := match(patterns, result)
	postprocessedlines := precompute(matchMap)
	serializedLines := out(postprocessedlines)
	jsonString := utils.MapToJsonString(serializedLines)
	return jsonString
}
