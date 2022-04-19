# Markdown To Notion
Export markdown into JSON.

Used to push markdown into Notion.

### Markdown

Each line can be categorized as follows:

- lineType: Either `paragraph_start_`, `paragraph_internal`, `paragraph_end`, `block` or `newline`
  - span types are included in the context of `paragraph_*`'s.
- results: The matchers that correspond to particular markdown elements e.g. 
```
{
  matcher: "inlinecode",
  indices: [0,34],
}
```
- content: The actual text content of the line, for an empty line we concatenate the line number with `idx` resulting in `idx:1`
