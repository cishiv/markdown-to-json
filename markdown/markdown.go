// Parse and convert a trivialised markdown spec to an opinionated JSON format
package markdown

import (
	"fmt"
	"strings"
)

type Pair[T, U any] struct {
	T any
	U any
}

type Block struct {
	content  string
	name     string
	children []Block
	emphasis Pair[int, int]
	strong   Pair[int, int]
	url      string
}

var blockpatterns = map[string]string{
	// These matchers _should_ work because we're matching blocks line-by-line
	"heading1":   "^#",
	"heading2":   "^##",
	"heading3":   "^###",
	"paragraph":  "DMAE", // doesnt match anything else
	"linebreak":  "\n",
	"blockquote": "^>",       // support only single level blocks (blocks can contain other elements)
	"codeblock":  "^```",     // this one needs depth
	"ul":         `\-|\*|\+`, // this one is debatable
	"ol":         `^[0-9]*\.`,
	"hr":         "^---",
}

var spanpatterns = map[string]string{
	// for em and strong, we need to count, but regex works
	"em":         "_.*_",
	"link":       `\[.*\]\(.*\)`,
	"strong":     `\*.*\*`,
	"inlinecode": "`.*`",
	"img":        "TODO", // TODO
}

var miscpatterns = map[string]string{
	"escape":   "TODO", // TODO
	"autolink": "TODO", // TODO
}

// func Parse(markdownString string) Block {
// 	return Block{
// 		name:     "body",
// 		content:  "",
// 		children: [
// 			Block{},
// 			Block{}
// 		],
// 	}
// }

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
		fmt.Println(fileSlice[ptr])
	}
	return ""
}

func patterns() {
	fmt.Println("block patterns")
	for name, pattern := range blockpatterns {
		if pattern == "DMAE" || pattern == "TODO" {
			continue
		}
		fmt.Println("Name:", name, "=>", "Pattern:", pattern)
	}

	fmt.Println("span patterns")
	for name, pattern := range spanpatterns {
		if pattern == "DMAE" || pattern == "TODO" {
			continue
		}
		fmt.Println("Name:", name, "=>", "Pattern:", pattern)
	}

	fmt.Println("misc patterns")
	for name, pattern := range miscpatterns {
		if pattern == "DMAE" || pattern == "TODO" {
			continue
		}
		fmt.Println("Name:", name, "=>", "Pattern:", pattern)
	}
}
