# Redate

A simple CLI tool that renames files by prefixing them with their last modified date (format: `yyyyMMddTHHmmss`), making file versioning and chronological sorting easy.

## Overview

Redate helps you automatically version files based on their last modified date. When you run the tool, it will rename files in a specified directory by adding a timestamp prefix. For example, a file named `document.txt` might be renamed to `20230615T143022_document.txt`.

This is particularly useful for:
- Tracking document versions without a formal version control system
- Ensuring chronological ordering in file listings
- Creating date-stamped backups of important files
- Managing shared documents across teams

## Installation

### Option 1: Download Binary (Recommended for most users)

Download the latest release for your platform from the [Releases](https://github.com/avanderw/redate/releases) page.

### Option 2: Build from Source

1. Ensure you have Go installed on your machine (version 1.16 or later recommended)
   ```
   go version
   ```

2. Clone the repository:
   ```
   git clone https://github.com/yourusername/redate.git
   cd redate
   ```

3. Build the executable:
   ```
   go build -o redate cmd/main.go
   ```

## Usage

Basic syntax:
```
redate [options] <pattern>
```

### Options

- `-r, --recursive`: Process files in subdirectories
- `-d, --dry-run`: Show what would be renamed without actually renaming
- `-f, --force`: Rename without prompting for confirmation

### Examples

**Basic usage:**
```
redate "*.txt"
```
This will find all `.txt` files in the current directory and prompt you to rename each one with a timestamp prefix.

**Recursive mode with specific file type:**
```
redate -r "*.docx"
```
This will find all `.docx` files in the current directory and all subdirectories.

**Sample output:**
```
$ redate "*.txt"
Found 3 files matching the pattern.
Rename "notes.txt" to "20230615T143022_notes.txt"? [y/N]: y
✓ Renamed to "20230615T143022_notes.txt"
Rename "todo.txt" to "20230614T092145_todo.txt"? [y/N]: n
⨯ Skipped "todo.txt"
Rename "report.txt" to "20230610T165513_report.txt"? [y/N]: y
✓ Renamed to "20230610T165513_report.txt"
```

## Troubleshooting

**Permission denied error:**
Make sure you have write permissions for the files you're trying to rename.

**No files found:**
Check that your pattern is correct and matches the files you expect.

## Why Use Redate?

While version control systems like Git are excellent for code, they can be cumbersome for non-code files or for users unfamiliar with VCS. Redate provides a simple way to keep track of file versions through timestamps, making it easy to identify the most recent version of a document.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

1. Fork the repository
2. Create a feature branch: `git checkout -b new-feature`
3. Commit your changes: `git commit -am 'Add new feature'`
4. Push to the branch: `git push origin new-feature`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.