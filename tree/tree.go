package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func BuildTree(path string) []string {
    var tree []string

    headerStyle, fileStyle, dirStyle, branchStyle := GetStyles()

    if AppConfig.ShowFolderIcon {
        tree = append(tree, headerStyle.Render(fmt.Sprintf("📂 %s", path)))
    }

    buildTree(path, "", &tree, fileStyle, dirStyle, branchStyle)
    return tree
}

func buildTree(path string, prefix string, tree *[]string, fileStyle, dirStyle, branchStyle lipgloss.Style) {
    files, err := os.ReadDir(path)
    if err != nil {
        *tree = append(*tree, fmt.Sprintf("Error reading directory: %v", err))
        return
    }

    sort.Slice(files, func(i, j int) bool {
        return files[i].Name() < files[j].Name()
    })

    var validFiles []os.DirEntry
    for _, file := range files {
        if AppConfig.HideDotFiles && strings.HasPrefix(file.Name(), ".") {
            continue
        }

        if AppConfig.DirsOnly && !file.IsDir() {
            continue
        }

        validFiles = append(validFiles, file)
    }

    for i, file := range validFiles {
        isLast := (i == len(validFiles) - 1)

        branch := branchStyle.Render("├── ")
        if isLast {
            branch = branchStyle.Render("└── ")
        }

        entry := branch + file.Name()
        if file.IsDir() {
            entry = branch + dirStyle.Render(file.Name())
        } else {
            entry = branch + fileStyle.Render(file.Name())
        }

        *tree = append(*tree, prefix+entry)

        if file.IsDir() {
            nextPrefix := prefix + branchStyle.Render("│   ")
            if isLast {
                nextPrefix = prefix + branchStyle.Render("    ")
            }
            buildTree(filepath.Join(path, file.Name()), nextPrefix, tree, fileStyle, dirStyle, branchStyle)
        }
    }

    if len(*tree) > 0 {
        *tree = append(*tree, "")
    }
}