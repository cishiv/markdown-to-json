package markdown

import "testing"

var unparsedMarkdown string = `
# This is a heading
## This is a second heading
- This is a list item
- This is also a list item
This should be a paragraph
---
**bold**
_italics_
[Link](https://link.com)
`

func TestMarkdownToProps(t *testing.T) {
	got := MarkdownToProps(unparsedMarkdown)
	want := `{
		"heading1": "",
		"heading2": "",
		"list": ["", ""],
		"paragraph": "",
		}`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
