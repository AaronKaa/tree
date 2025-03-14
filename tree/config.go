package tree

import (
    "bufio"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

type Config struct {
    ColorFile       string
    ColorDir        string
    ColorBranch     string
    HideDotFiles    bool
    Summarize       bool
    ShowFolderIcon  bool
    DirsOnly        bool
    FolderIcon      string
}

var AppConfig = Config{
    ColorFile:      "#FF0883",
    ColorDir:       "#83FF08",
    ColorBranch:    "#B6B6B6",
    HideDotFiles:   false,
    Summarize:      false,
    ShowFolderIcon: true,
    DirsOnly:       false,
    FolderIcon:     "📂",
}

var hexColorPattern = regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)

func validateHexColor(color, fallback string) string {
    color = strings.TrimSpace(strings.Trim(color, `"'`))

    if strings.Count(color, "#") > 1 {
        color = strings.Split(color, "#")[0]
        color = strings.TrimSpace(color)
    }

    if hexColorPattern.MatchString(color) {
        return color
    }

    return fallback
}

func LoadConfig() {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return
    }

    configPath := filepath.Join(homeDir, ".tree")
    file, err := os.Open(configPath)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }

        key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

        switch key {
        case "color_file":
            AppConfig.ColorFile = validateHexColor(value, AppConfig.ColorFile)
        case "color_dir":
            AppConfig.ColorDir = validateHexColor(value, AppConfig.ColorDir)
        case "color_branch":
            AppConfig.ColorBranch = validateHexColor(value, AppConfig.ColorBranch)
        case "hide_dotfiles":
            AppConfig.HideDotFiles = (value == "true")
        case "summarize":
            AppConfig.Summarize = (value == "true")
        case "show_folder_icon":
            AppConfig.ShowFolderIcon = (value == "true")
        case "dirs_only":
            AppConfig.DirsOnly = (value == "true")
        case "folder_icon":
            AppConfig.FolderIcon = value
        }
    }

    _ = scanner.Err()
}