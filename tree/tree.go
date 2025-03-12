package tree

import (
    "fmt"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

func buildTree(path, prefix string, tree *[]string) {
    files, err := os.ReadDir(path)
    if err != nil {
        *tree = append(*tree, fmt.Sprintf("Error: %v", err))
        return
    }

    // Filter and sort files
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
    sort.Slice(validFiles, func(i, j int) bool {
        return validFiles[i].Name() < validFiles[j].Name()
    })

    // Get styles
    _, fileStyle, dirStyle, branchStyle := GetStyles()

    for i, file := range validFiles {
        isLast := i == len(validFiles)-1

        // Determine branch symbol
        branch := "├── "
        if isLast {
            branch = "└── "
        }
        branch = branchStyle.Render(branch)

        // Render entry
        entry := file.Name()
        if file.IsDir() {
            entry = dirStyle.Render(entry)
        } else {
            entry = fileStyle.Render(entry)
        }
        *tree = append(*tree, prefix+branch+entry)

        // Recursively build tree for directories
        if file.IsDir() {
            nextPrefix := prefix + branchStyle.Render("│   ")
            if isLast {
                nextPrefix = prefix + "    "
            }
            buildTree(filepath.Join(path, file.Name()), nextPrefix, tree)
        }
    }
}

func BuildTree(path string) []string {
    var tree []string

    headerStyle, _, _, _ := GetStyles()
    tree = append(tree, headerStyle.Render(fmt.Sprintf("📂 %s", path)))

    buildTree(path, "", &tree)
    return tree
}