package notion

import "obsidian-to-notion/markdown"

type NotionBlock struct {
	typeName string
	content  string
}

var NotionBlockSpec = map[markdown.MdPattern]string{
	{
		Name:    "h1",
		Pattern: "#",
	}: "heading_1",
	{
		Name:    "h2",
		Pattern: "##",
	}: "heading_2",
	{
		Name:    "list",
		Pattern: "-",
	}: "heading_1",
}
