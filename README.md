# Markdown To Notion
Export markdown into JSON.

`go get github.com/cishiv/markdown-to-json/v2`

Given a markdown file:

`/markdown/examples/small-test.md`
```
# h1
## h2
### h3
paragraph `inlinetest`
paragraph
paragraph

newparagraph
newparagraph
newparagraph
```

We produce the following structured JSON:
```
{
  "0": {
    "content": "# h1",
    "results": [
      {
        "matcher": "heading1",
        "occurences": [
          {
            "firstIdx": 0,
            "secondIdx": 1
          }
        ]
      }
    ],
    "type": "block"
  },
  "1": {
    "content": "## h2",
    "results": [
      {
        "matcher": "heading2",
        "occurences": [
          {
            "firstIdx": 0,
            "secondIdx": 2
          }
        ]
      }
    ],
    "type": "block"
  },
  "2": {
    "content": "### h3",
    "results": [
      {
        "matcher": "heading3",
        "occurences": [
          {
            "firstIdx": 0,
            "secondIdx": 3
          }
        ]
      }
    ],
    "type": "block"
  },
  "3": {
    "content": "paragraph `inlinetest`",
    "results": [
      {
        "matcher": "inlinecode",
        "occurences": [
          {
            "firstIdx": 10,
            "secondIdx": 22
          }
        ]
      }
    ],
    "type": "paragraph_start"
  },
  "4": {
    "content": "paragraph",
    "results": [
      {
        "matcher": "paragraph",
        "occurences": null
      }
    ],
    "type": "paragraph_internal"
  },
  "5": {
    "content": "paragraph",
    "results": [
      {
        "matcher": "paragraph",
        "occurences": null
      }
    ],
    "type": "paragraph_end"
  },
  "6": {
    "content": "idx:6",
    "results": [
      {
        "matcher": "newline",
        "occurences": null
      }
    ],
    "type": "newline"
  },
  "7": {
    "content": "newparagraph",
    "results": [
      {
        "matcher": "paragraph",
        "occurences": null
      }
    ],
    "type": "paragraph_start"
  },
  "8": {
    "content": "newparagraph",
    "results": [
      {
        "matcher": "paragraph",
        "occurences": null
      }
    ],
    "type": "paragraph_internal"
  },
  "9": {
    "content": "newparagraph",
    "results": [
      {
        "matcher": "paragraph",
        "occurences": null
      }
    ],
    "type": "paragraph_end"
  }
}
```

You can reproduce this result by running `./test.sh` or:

```
cd markdown
go test
```