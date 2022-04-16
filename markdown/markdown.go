// Parse and convert a trivialised markdown spec to an opinionated JSON format
package markdown

import (
	"fmt"
	"regexp"
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

// paragraph is fallback
var blockpatterns = map[string]string{
	// These matchers _should_ work because we're matching blocks line-by-line
	"heading1":   "^#",
	"heading2":   "^##",
	"heading3":   "^###",
	"linebreak":  "\n",
	"blockquote": "^>",       // support only single level blocks (blocks can contain other elements).
	"codeblock":  "^```",     // this one needs depth.
	"ul":         `\-|\*|\+`, // this one is debatable.
	"ol":         `^[0-9]*\.`,
	"hr":         "^---",
}

// excludes images.
var spanpatterns = map[string]string{
	// for em and strong, we need to count, but regex works.
	"em":         "_.*_",
	"link":       `\[.*\]\(.*\)`,
	"strong":     `\*.*\*`,
	"inlinecode": "`.*`",
	"img":        "TODO", // TODO
}

// do not use - just for context / remembering that these might be useful things.
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

func compilePatterns() map[string]regexp.Regexp {
	fmt.Println("compile block patterns")
	patterns := make(map[string]regexp.Regexp)
	for name, pattern := range blockpatterns {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		patterns[pattern] = *compiled.Copy()
		fmt.Println("Complied Name:", name, "=>", "Pattern:", pattern)
	}

	fmt.Println("compile span patterns")
	for name, pattern := range spanpatterns {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		patterns[pattern] = *compiled.Copy()
		fmt.Println("Compiled Name:", name, "=>", "Pattern:", pattern)
	}
	return patterns
}