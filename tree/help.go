package tree

import "fmt"

// PrintHelp prints the CLI help message
func PrintHelp() {
    fmt.Println(`
Tree - A dir tree command

Usage:
  tree [path] [options]

Options:
  --help, --h            Show help information
  --summarise, --s       Summarize directory contents
  --no-summarise, --ns   Show all files without summarizing
  --dot-files, --df      Show dot files (e.g., .git, .env)
  --no-dot-files, --ndf  Hide dot files
  --dirs-only, --do      Show only directories (hide files)

Examples:
  tree                # Show tree for current directory
  tree ~/projects     # Show tree for specific path
  tree --s --ndf      # Summarize and hide dot files
  tree --df --do      # Show only directories and dot files`)
}