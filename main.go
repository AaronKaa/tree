package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/AaronKaa/tree/tree"
)

func printHelp() {
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
  tree --df --do      # Show only directories and dot files
`)
}

func main() {
    var path string

    tree.LoadConfig()

    if len(os.Args) > 1 {
        for _, arg := range os.Args[1:] {
            switch arg {
            case "--help", "--h":
                printHelp()
                return
            case "--summarise", "--s":
                tree.AppConfig.Summarize = true
            case "--no-summarise", "--ns":
                tree.AppConfig.Summarize = false
            case "--dot-files", "--df":
                tree.AppConfig.HideDotFiles = false
            case "--no-dot-files", "--ndf":
                tree.AppConfig.HideDotFiles = true
            case "--dirs-only", "--do":
                tree.AppConfig.DirsOnly = true
            default:
                if path == "" {
                    path = arg
                }
            }
        }
    }

    if path == "" {
        path = "."
    }

    m := tree.NewModel(path)
    p := tea.NewProgram(m, tea.WithOutput(os.Stdout))

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}