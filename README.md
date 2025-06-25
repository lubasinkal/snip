# snip — A Terminal Code Snippet Manager

A superfast command-line tool to save, search, view, and reuse your code snippets — all from your terminal. Think fzf + ripgrep but for your personal knowledge bank of code.

## ✨ Features

- 🚀 **Lightning fast** - Built in Go for speed
- 💾 **Local storage** - Your snippets stay on your machine (SQLite database)
- 🔍 **Powerful search** - Search by title, content, or tags
- 🏷️ **Tag support** - Organize snippets with multiple tags
- 📋 **Clipboard integration** - Copy snippets directly to clipboard
- ✏️ **Edit in place** - Open snippets in your favorite editor
- 🎯 **Simple CLI** - Intuitive commands that just work
- 📊 **Statistics** - Beautiful analytics about your snippets
- 📤 **Export/Import** - Backup and restore your collection
- 🎨 **Beautiful UI** - Powered by Charm's Lipgloss & Huh
- 🎪 **Interactive mode** - Guided snippet creation with forms

## 🚀 Quick Start

```bash
# Initialize (optional - creates ~/.snipdb directory)
snip init

# Save a snippet from stdin
echo 'fmt.Println("Hello, World!")' | snip save "Go Hello World" --tags=go,example

# Save interactively with guided form
snip save-interactive

# Save from a file
snip save "My bash script" --tags=bash,utility < script.sh

# List all snippets
snip list

# Search snippets
snip search "hello"
snip search "function" --tag=javascript

# View a snippet
snip cat 1

# Copy to clipboard
snip copy 1

# Edit a snippet
snip edit 1

# Delete a snippet
snip delete 1

# View statistics
snip stats

# Export snippets
snip export --format=json --output=backup.json

# Import snippets
snip import backup.json
```

## 📖 Command Reference

### `snip save` - Save a snippet
```bash
# Save from stdin
echo 'your code here' | snip save "Description"

# Save with tags
echo 'SELECT * FROM users;' | snip save "Get all users" --tags=sql,database

# Save from file
snip save "My config" --tags=config < ~/.bashrc
```

### `snip list` - List all snippets
```bash
snip list
```
Shows all snippets with ID, title, tags, and creation time.

### `snip search` - Search snippets
```bash
# Search by content or title
snip search "function"

# Search with tag filter
snip search "http" --tag=javascript

# Search by tag only
snip search "" --tag=python
```

### `snip cat` - View snippet content
```bash
# Print to stdout (perfect for piping)
snip cat 1

# Pipe to other commands
snip cat 1 | grep "TODO"
```

### `snip copy` - Copy to clipboard
```bash
snip copy 1
```

### `snip edit` - Edit snippet
```bash
snip edit 1
```
Opens the snippet in your default editor (`$EDITOR` environment variable).

### `snip delete` - Delete snippet
```bash
# With confirmation
snip delete 1

# Skip confirmation
snip delete 1 --force
```

### `snip save-interactive` - Interactive snippet creation
```bash
snip save-interactive
```
Launch an interactive form to save a code snippet with guided prompts for title, content, language, and tags.

### `snip stats` - View statistics
```bash
snip stats
```
Display beautiful analytics about your snippet collection including counts, popular tags, and recent activity.

### `snip export` - Export snippets
```bash
# Export to JSON (default)
snip export

# Export to specific format and file
snip export --format=markdown --output=my_snippets.md

# Available formats: json, markdown, text
snip export --format=json --output=backup.json
```

### `snip import` - Import snippets
```bash
# Import from file
snip import backup.json

# Import with confirmation skip
snip import --yes backup.json
```

### `snip version` - Show version information
```bash
snip version
```
Display version, features, and credits information.

### `snip init` - Initialize configuration
```bash
snip init
```
Creates `~/.snipdb/` directory and shows setup information.

## 🛠️ Installation

### From Source
```bash
git clone https://github.com/lubasinkal/snip.git
cd snip
go build -o snip .
```

### Binary Release
Download the latest binary from the [releases page](https://github.com/lubasinkal/snip/releases).

## 🏗️ Architecture

- **Storage**: SQLite database at `~/.snipdb/snippets.db`
- **Search**: Full-text search across titles, content, and tags
- **Clipboard**: Cross-platform clipboard support via `github.com/atotto/clipboard`
- **Editor**: Respects `$EDITOR` environment variable with sensible defaults
- **UI Framework**: Beautiful terminal interfaces powered by Charm's Lipgloss and Huh
- **Interactive Forms**: Rich form-based input with validation and language selection

## 💡 Examples

### Save common code patterns
```bash
# Go HTTP handler
echo 'func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}' | snip save "Basic HTTP handler" --tags=go,http,web

# Python list comprehension
echo '[x**2 for x in range(10) if x % 2 == 0]' | snip save "Even squares" --tags=python,list-comprehension

# SQL query
echo 'SELECT u.name, COUNT(o.id) as order_count
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
GROUP BY u.id;' | snip save "User order counts" --tags=sql,join,aggregate
```

### Interactive snippet creation
```bash
# Use the guided form for easy snippet creation
snip save-interactive
# Follow the prompts to enter title, select language, add content, and tags
```

### Search and reuse
```bash
# Find all Python snippets
snip search "" --tag=python

# Find HTTP-related code
snip search "http"

# Copy a snippet and use it
snip copy 5
# Paste in your editor (Ctrl+V)
```

### Backup and restore
```bash
# Create a backup of all snippets
snip export --format=json --output=my_backup.json

# Restore from backup
snip import my_backup.json

# Export to markdown for documentation
snip export --format=markdown --output=snippets_doc.md
```

## 🎨 Enhanced User Experience

### Beautiful Terminal UI
- **Rich formatting** with colors and styling powered by Lipgloss
- **Interactive forms** with validation and smart defaults
- **Progress indicators** and status messages
- **Syntax highlighting** for code snippets
- **Responsive tables** that adapt to terminal width

### Smart Workflows
- **Language detection** - Interactive mode suggests appropriate tags
- **Validation** - Forms prevent empty content and invalid input
- **Confirmation prompts** - Safe operations with clear feedback
- **Preview modes** - See what you're importing before committing

## 🔧 Configuration

### Environment Variables
- `EDITOR` - Your preferred text editor for `snip edit`
- `SNIP_DB_PATH` - Custom database location (default: `~/.snipdb/snippets.db`)

### Default Editors
If `$EDITOR` is not set, snip will try to use (in order):
1. `code` (VS Code)
2. `nano`
3. `vim`
4. `notepad` (Windows)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

Built with amazing open source libraries:

- **[Cobra](https://github.com/spf13/cobra)** - CLI framework by spf13
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Terminal styling by Charm
- **[Huh](https://github.com/charmbracelet/huh)** - Terminal forms by Charm
- **[modernc.org/sqlite](https://gitlab.com/cznic/sqlite)** - Embedded SQLite database
- **[atotto/clipboard](https://github.com/atotto/clipboard)** - Cross-platform clipboard support