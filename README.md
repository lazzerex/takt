# Takt Auto Clicker

A safe and minimal auto clicker built with **Go** and **Wails** for Windows

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
