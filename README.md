snip — A Terminal Code Snippet Manager

A superfast command-line tool to save, search, view, and reuse your code snippets — all from your terminal. Think fzf + ripgrep but for your personal knowledge bank of code.
🚀 Core User Flow (Minimal but Powerful)

# Save a snippet from file or stdin
snip save "List files in Go" < listfiles.go

# Save with tags
snip save "Basic HTTP server" --tag=go,web < server.go

# List snippets
snip list

# Search snippets by text or tag
snip search "http"
snip search --tag=go

# View or copy a snippet
snip cat 3
snip copy 3

# Delete or edit a snippet
snip delete 3
snip edit 3

🧱 How It Will Work (Architecture Overview)
1. Storage:

Each snippet is saved as a file + metadata:

    Store in ~/.snipdb/ (or configurable path)


{
  "id": 3,
  "title": "Basic HTTP server",
  "tags": ["go", "web"],
  "created_at": "2025-06-25T12:30:00Z",
  "filepath": "/home/lubasi/.snipdb/snippets/3.go"
}

2. Indexing / Search

Use simple substring or fuzzy match (github.com/lithammer/fuzzysearch or fzf integration) to search through:

    Titles

    Tags

    Content

3. Clipboard Integration

    On Linux/macOS: use xclip, pbcopy, or wl-copy

    On Windows: Go bindings for clipboard (e.g. github.com/atotto/clipboard)

4. Command-line UX:

🧪 MVP Feature Set
Feature	Description
snip save	Save stdin or a file as a snippet with title + optional tags
snip list	List all snippets by ID, title, and tags
snip search	Search snippets by title, tags, or content
snip cat	Print snippet to stdout
snip copy	Copy snippet content to clipboard
snip delete	Remove snippet
snip edit	Open snippet in $EDITOR
snip init	Setup config or storage dir (optional)