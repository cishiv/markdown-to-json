package markdown

type Classification string

const (
	PARAGRAPH_START Classification = "paragraph_start"
)

const NEWLINE = "newline"
const PARAGRAPH = "paragraph"

const BLOCK_PATTERN = "block"
const SPAN_PATTERN = "span"

// paragraph is fallback
var blockpatterns = map[string]string{
	// These matchers _should_ work because we're matching blocks line-by-line
	// recurrence capture doesn't bubble up the way I expect, this needs to be handled programatically
	// if something is an ###, it can't be ## or #
	"heading1":   "^#{1}",
	"heading2":   "^#{2}",
	"heading3":   "^#{3}",
	"linebreak":  "\n",
	"blockquote": "^>",          // support only single level blocks (blocks can contain other elements).
	"codeblock":  "^```",        // this one needs depth.
	"ul":         `^\-|^\*|^\+`, // this one is debatable.
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
