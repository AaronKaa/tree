# Tree

tree is a CLI tool written in Go that displays a file tree structure in the terminal. There are plenty of "tree" apps but this one is designed to make easily pasteable dir trees for software documentation / tutorials.

## INSTALLATION

1. Clone the repository:
   git clone https://github.com/AaronKaa/tree.git
   cd tree

2. Build the project:
   go build -o tree .

3. Move the binary to your path:
   mv tree /usr/local/bin/

## USAGE

### Basic Command:

To display the file tree for the current directory:

    tree

### Specify Path:

To display the file tree for a specific directory:

    tree ~/projects

### Configuration With `.tree` File

You can create a `.tree` config file in your home directory to control the styling and behaviour:

    touch ~/.tree

### Example `.tree` Configuration:

```
# Colors (hex codes)
color_file = #FF5733
color_dir = #4CAF50
color_branch = #FFC107

# Behavior
hide_dotfiles = true
summarize = false
show_folder_icon = true

```

### Available Config Options

| Option           | Description                  | Example      |
| ---------------- | ---------------------------- | ------------ |
| color_file       | Color for files              | #FF5733      |
| color_dir        | Color for directories        | #4CAF50      |
| color_branch     | Color for branch lines       | #FFC107      |
| hide_dotfiles    | Hide dot files               | true / false |
| summarize        | Summarize directory contents | true / false |
| show_folder_icon | Show 📂 icon at the top      | true / false |

### Runtime Options

| Switch         | Shortcut | Description                        |
| -------------- | -------- | ---------------------------------- |
| --summarise    | --s      | Summarize directory contents       |
| --no-summarise | --ns     | Show all files without summarizing |
| --dot-files    | --df     | Show dot files (e.g., .git, .env)  |
| --no-dot-files | --ndf    | Hide dot files                     |
| --dirs-only    | --do     | Show only directories (hide files) |

## Examples:

Summarise directory contents

    `tree --summarise`

Hide dot files:

    `tree --no-dot-files`

Summarise and hide dot files:

    `tree --s --ndf`

Show dot files and avoid summarising:

    `tree --df --ns`

Override config and display specific path:

    `tree ~/projects --s --ndf`

## EXAMPLES

## Example with default config:

    tree

```
📂 .
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   │   ├── main.go
│   │   │   ├── handlers/
│   │   │   └── routes/
│   │   └── subscriber/
│   │       ├── main.go
│   │       └── event_handler.go
├── internal/
│   ├── models/
│   │   ├── user.go
│   │   ├── link.go
│   │   ├── click.go
```

## Example with summarising and hidden dot files:

    tree --s --ndf

```
📂 .
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   └── subscriber/
├── internal/
│   ├── models/
```

## Example with directories only

    tree --do

```
📂 .
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   └── subscriber/
├── internal/
│   ├── models/
```

## LICENSE

This project is licensed under the MIT License. See the LICENSE file for more information.
