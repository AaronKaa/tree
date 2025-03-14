package tree

import (
    "fmt"
    "os"
    "path/filepath"
    "sort"
    "strings"

    liptree "github.com/charmbracelet/lipgloss/tree"
)

func buildTree(node *liptree.Tree, path string) (int, int) {
    files, err := os.ReadDir(path)
    if err != nil {
        return 0, 0
    }

    sort.Slice(files, func(i, j int) bool {
        return files[i].Name() < files[j].Name()
    })

    _, fileStyle, dirStyle, _ := GetStyles()

    fileCount := 0
    dirCount := 0

    for _, file := range files {
        if AppConfig.HideDotFiles && strings.HasPrefix(file.Name(), ".") {
            continue
        }

        if AppConfig.DirsOnly && !file.IsDir() {
            continue
        }

        if file.IsDir() {
            dirCount++
            name := file.Name()

            child := liptree.New().
                Root(dirStyle.Render(name))

            childFiles, childDirs := buildTree(child, filepath.Join(path, file.Name()))

            if !AppConfig.Summarize {
                node.Child(child)
            }

            fileCount += childFiles
            dirCount += childDirs
        } else {
            fileCount++
            if !AppConfig.DirsOnly {
                node.Child(
                    liptree.New().
                        Root(fileStyle.Render(file.Name())),
                )
            }
        }
    }

    return fileCount, dirCount
}

func BuildTree(path string) *liptree.Tree {
    _, _, _, branchStyle := GetStyles()

    if path == "." {
        absPath, err := filepath.Abs(path)
        if err == nil {
            path = absPath
        }
    }

    rootName := filepath.Base(path)

    if AppConfig.ShowFolderIcon {
        rootName = fmt.Sprintf("%s %s", AppConfig.FolderIcon, rootName)
    }

    root := liptree.New().
        Root(branchStyle.Render(rootName))

    fileCount, dirCount := buildTree(root, path)

    if AppConfig.Summarize {
        summary := fmt.Sprintf("(%d files, %d directories)", fileCount, dirCount)
        rootName = fmt.Sprintf("%s %s", rootName, summary)
        root = liptree.New().Root(branchStyle.Render(rootName))
    }

    return root
}