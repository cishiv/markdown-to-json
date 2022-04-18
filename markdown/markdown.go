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

func copyAndAddLineTypeToLinkedLine(tar LinkedLine, lineType string) LinkedLine {
	return LinkedLine{
		resultStrings:   tar.resultStrings,
		unparsedResults: tar.unparsedResults,
		safe:            tar.safe,
		lineType:        lineType,
		content:         tar.content,
	}
}
