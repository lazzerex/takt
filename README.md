<h1 align="center">Takt Auto Clicker</h1>

<p align="center">
  <img width="427" height="653" alt="takt"
       src="https://github.com/user-attachments/assets/36c90eb3-d8a7-4af6-a97b-66d342482b8b" />
</p>

<p align="center">
  <strong>A safe and minimal auto clicker built with <b>Go</b> and <b>Wails</b> for Windows</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white"/>
  <img src="https://img.shields.io/badge/Wails-DF0000?style=flat&logo=go&logoColor=white"/>
  <img src="https://img.shields.io/badge/Windows-0078D6?style=flat&logo=windows&logoColor=white"/>
  <img src="https://img.shields.io/badge/Desktop%20App-4B32C3?style=flat&logo=electron&logoColor=white"/>
  <img src="https://img.shields.io/badge/Lightweight-2ECC71?style=flat"/>
  <img src="https://img.shields.io/badge/Open%20Source-000000?style=flat&logo=github&logoColor=white"/>
</p>

<p align="center">
  <!-- Stars -->
  <img src="https://img.shields.io/github/stars/lazzerex/takt?style=flat&logo=github&logoColor=white"/>
  <!-- Forks -->
  <img src="https://img.shields.io/github/forks/lazzerex/takt?style=flat&logo=github&logoColor=white"/>
  <!-- Open Issues -->
  <img src="https://img.shields.io/github/issues/lazzerex/takt?style=flat&color=red"/>
  <!-- Pull Requests -->
  <img src="https://img.shields.io/github/issues-pr/lazzerex/takt?style=flat&color=yellow"/>
  <!-- Last Commit -->
  <img src="https://img.shields.io/github/last-commit/lazzerex/takt?style=flat"/>
  <!-- Repo Size -->
  <img src="https://img.shields.io/github/repo-size/lazzerex/takt?style=flat"/>
</p>



## Features

- **Native Performance**: Built with Go for blazing-fast execution
- **Modern UI**: Clean, gradient-based design with smooth animations
- **Customizable**: Adjust click intervals, mouse buttons, and click types
- **Global Hotkey**: Ctrl+S toggle works even when minimized
- **Responsive Layout**: Adapts to different window sizes

## Tech Stack

- **Backend**: Go (Golang)
- **GUI Framework**: Wails v2
- **Frontend**: Vanilla JavaScript with modern CSS

## Installation

### Prerequisites

1. **Go 1.21+**
   ```bash
   # Download from https://golang.org/dl/
   ```

2. **Wails CLI**
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

3. **Platform-specific dependencies**:
   - **Windows**: No additional dependencies
   - **macOS**: Xcode Command Line Tools
   - **Linux**: `libgtk-3-dev libwebkit2gtk-4.0-dev`

### Build & Run

1. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

2. **Development mode**
   ```bash
   wails dev
   ```

3. **Build production binary**
   ```bash
   wails build
   ```

The compiled executable will be in `build/bin/`

## Usage

1. **Launch Takt**
   - Run the executable or use `wails dev`

2. **Configure Settings**
   - **Interval**: Set click interval in seconds (supports decimals)
   - **Mouse Button**: Choose Left, Right, or Middle
   - **Click Type**: Select Single or Double click

3. **Start Clicking**
   - Click "START" button or press **Ctrl+S**
   - Press **Ctrl+S** again or click "STOP" to stop

## Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `Ctrl+S` | Toggle auto clicker on/off (works globally, even when app is minimized) |

## Project Structure

```
takt/
├── main.go              # Application entry point
├── app.go               # Core auto clicker logic
├── go.mod               # Go dependencies
├── frontend/
│   └── dist/
│       ├── index.html   # UI layout
│       └── app.js       # Frontend logic
└── README.md
```

## Configuration

All settings are configured through the UI:

- **Click Interval**: 0.01 - unlimited seconds
- **Mouse Buttons**: Left, Right, Middle
- **Click Types**: Single, Double

## Responsible Use

This tool is designed for legitimate automation tasks. Please use responsibly:

- Testing applications
- Repetitive tasks in appropriate software
- Accessibility assistance

Do not use for violating terms of service or bypassing security measures.

## License

MIT License - Feel free to use and modify as needed.

## Contributing

Contributions welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests

## Acknowledgments

- [Wails](https://wails.io/) - Go framework for desktop apps
- Windows API for native mouse control and global hotkeys

---

*Fast. Clean. Efficient.*
