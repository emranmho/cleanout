# Cleanout CLI 🧹

A cross-platform command-line utility for efficient file cleanup based on file age. Cleanout helps you automate the removal of temporary files, cache, logs, and other unnecessary files from your system.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.24+-00ADD8.svg)

## 🌟 Features

- **Automated Cleanup**: Scan directories and remove files older than a specified age
- **Dry Run Mode**: Preview deletions without actually removing files
- **Multiple Commands**:
  - `clean`: Main cleanup command for removing old files
  - `clean-logs`: Specialized command for cleaning the tool's own log files
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Detailed Logging**: Comprehensive logs of all operations
- **User-Friendly Output**: Clean, emoji-rich console output with operation summaries
- **Safety Measures**: Directory checks and access verification

## 📋 Project Structure

```
cleanout/
├── cmd/                     # Command implementations
│   ├── clean.go             # Main cleaning functionality
│   ├── clean-log.go         # Log cleanup functionality
│   └── root.go              # Root command definition
├── internal/                # Internal packages
│   ├── cleaner.go           # Core file cleaning logic
│   └── util.go              # Utilities and logging
├── logs/                    # Generated logs (created at runtime)
├── build.ps1                # PowerShell build script for cross-platform
├── Makefile                 # Unix-style build commands
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── main.go                  # Application entry point
└── README.md                # Project documentation
```

## 💻 Installation

### Using Go

```bash
go install github.com/emranmho/cleanout@latest
```

### Manual Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/emranmho/cleanout.git
   cd cleanout
   ```

2. Build the binary:
   ```bash
   # Using make (Unix/Linux/macOS)
   make build
   
   # Or using go directly
   go build -o cleanout .
   
   # For Windows (PowerShell)
   .\build.ps1
   ```

3. Move the binary to your PATH (optional):
   ```bash
   # Example for Linux/macOS
   sudo mv cleanout /usr/local/bin/
   ```

## 🚀 Usage

### Basic Usage

```bash
# Clean files older than 7 days in system temp directory
cleanout clean

# Clean files with custom parameters
cleanout clean --path "/path/to/directory" --days 30 --verbose
```

### Command Options

#### `clean` Command

```bash
# Preview deletion (dry-run mode)
cleanout clean --dry-run

# Clean a specific directory
cleanout clean --path "C:\Users\Username\Downloads"

# Delete files older than 14 days with detailed output
cleanout clean --days 14 --verbose
```

#### `clean-logs` Command

```bash
# Delete logs older than 7 days (default)
cleanout clean-logs

# Delete logs older than 14 days
cleanout clean-logs --days 14
```

### Available Flags

**For `clean` command:**
- `--path`: Directory to scan (default: system temp directory)
- `--days`: Delete files older than N days (default: 7)
- `--dry-run`: Preview deletions without actually deleting
- `--verbose`: Show detailed progress and file operations

**For `clean-logs` command:**
- `--days`: Delete log files older than N days (default: 7)

## 📊 Logs

Cleanout automatically generates detailed logs of all operations in the `logs` directory:

- Log files are named with a timestamp: `cleanup_YYYYMMDD_HHMMSS.log`
- Dry-run logs are marked with `cleanup_dry_run_YYYYMMDD_HHMMSS.log`
- Logs include operation details, timestamps, and summary statistics
- Use the `clean-logs` command to manage older log files

## 🛠️ Developer Setup

### Prerequisites

- Go 1.24 or higher
- Git

### Setup Development Environment

1. Clone the repository:
   ```bash
   git clone https://github.com/emranmho/cleanout.git
   cd cleanout
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

### Building for Multiple Platforms

#### Using PowerShell (Windows):

```powershell
.\build.ps1
```

This will generate:
- `cleanout.exe` for Windows
- `cleanout` for Linux
- `cleanout-mac` for macOS

#### Using Makefile (Unix/Linux/macOS):

```bash
make build   # Build for current platform
make clean   # Clean build artifacts
```

### Running Tests

```bash
go test ./...
```

## 🗺️ Roadmap

- [ ] Add automated tests
- [ ] Implement file pattern matching for more targeted cleanup
- [ ] Add scheduler integration for automated periodic cleanup
- [ ] Create interactive terminal UI
- [ ] Add support for cleaning based on file types
- [ ] Implement file size-based cleaning options
- [ ] Add file recovery option for recently deleted files
- [ ] Create system service integration for background operation
- [ ] Add configuration file support
- [ ] Implement cloud storage cleanup options

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
