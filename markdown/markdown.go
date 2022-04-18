// Parse and convert a trivialised markdown spec to an opinionated JSON format
package markdown

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
	return ""
}

func FromMarkdownStringToJsonString(markdownString string) string {
	return ""
}

func FromMarkdownFileToJsonFile(markdownFilePath string, jsonFilePath string) {
}

func FromMarkdownStringToJsonFile(markdownString string, jsonFilePath string) {

}

// To Markdown

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
