package main

import (
	"context"
	"fmt"
	goruntime "runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx        context.Context
	isRunning  bool
	mutex      sync.Mutex
	stopChan   chan bool
	lastConfig ClickerConfig
	hotkeyQuit chan bool
}

func NewApp() *App {
	return &App{
		isRunning:  false,
		stopChan:   make(chan bool),
		hotkeyQuit: make(chan bool),
		lastConfig: ClickerConfig{
			Interval:  1.0,
			Button:    "left",
			ClickType: "single",
		},
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.registerGlobalHotkey()
}

func (a *App) shutdown(ctx context.Context) {
	a.StopClicking()
	close(a.hotkeyQuit)
}

type ClickerConfig struct {
	Interval  float64 `json:"interval"`
	Button    string  `json:"button"`
	ClickType string  `json:"clickType"`
}

func (a *App) StartClicking(config ClickerConfig) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.isRunning {
		return fmt.Errorf("auto clicker is already running")
	}

	if config.Interval <= 0 {
		return fmt.Errorf("interval must be greater than 0")
	}

	a.lastConfig = config
	a.isRunning = true
	a.stopChan = make(chan bool)

	go a.clickLoop(config)

	runtime.EventsEmit(a.ctx, "status_changed", "running")
	return nil
}

func (a *App) StopClicking() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.isRunning {
		return
	}

	a.isRunning = false
	close(a.stopChan)

	runtime.EventsEmit(a.ctx, "status_changed", "idle")
}

func (a *App) IsRunning() bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.isRunning
}

func (a *App) clickLoop(config ClickerConfig) {
	ticker := time.NewTicker(time.Duration(config.Interval * float64(time.Second)))
	defer ticker.Stop()

	for {
		select {
		case <-a.stopChan:
			return
		case <-ticker.C:
			a.performClick(config.Button, config.ClickType)
		}
	}
}

const (
	MOUSEEVENTF_LEFTDOWN   = 0x0002
	MOUSEEVENTF_LEFTUP     = 0x0004
	MOUSEEVENTF_RIGHTDOWN  = 0x0008
	MOUSEEVENTF_RIGHTUP    = 0x0010
	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP   = 0x0040
	WM_HOTKEY              = 0x0312
	VK_S                   = 0x53
	MOD_CONTROL            = 0x0002
	MOD_NOREPEAT           = 0x4000
	HOTKEY_ID              = 1
	PM_REMOVE              = 0x0001
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procMouseEvent       = user32.NewProc("mouse_event")
	procRegisterHotKey   = user32.NewProc("RegisterHotKey")
	procUnregisterHotKey = user32.NewProc("UnregisterHotKey")
	procPeekMessage      = user32.NewProc("PeekMessageW")
	procTranslateMessage = user32.NewProc("TranslateMessage")
	procDispatchMessage  = user32.NewProc("DispatchMessageW")
)

func (a *App) performClick(button string, clickType string) {
	var downFlag, upFlag uintptr

	switch button {
	case "right":
		downFlag = MOUSEEVENTF_RIGHTDOWN
		upFlag = MOUSEEVENTF_RIGHTUP
	case "middle":
		downFlag = MOUSEEVENTF_MIDDLEDOWN
		upFlag = MOUSEEVENTF_MIDDLEUP
	default:
		downFlag = MOUSEEVENTF_LEFTDOWN
		upFlag = MOUSEEVENTF_LEFTUP
	}

	mouseClick(downFlag, upFlag)

	if clickType == "double" {
		time.Sleep(50 * time.Millisecond)
		mouseClick(downFlag, upFlag)
	}
}

func mouseClick(downFlag, upFlag uintptr) {
	procMouseEvent.Call(downFlag, 0, 0, 0, 0)
	time.Sleep(10 * time.Millisecond)
	procMouseEvent.Call(upFlag, 0, 0, 0, 0)
}

// registerGlobalHotkey registers Ctrl+S as a global hotkey.
// Must lock OS thread since Windows hotkey messages are thread-specific.
func (a *App) registerGlobalHotkey() {
	goruntime.LockOSThread()
	defer goruntime.UnlockOSThread()

	ret, _, _ := procRegisterHotKey.Call(0, HOTKEY_ID, MOD_CONTROL|MOD_NOREPEAT, VK_S)
	if ret == 0 {
		fmt.Println("Failed to register Ctrl+S hotkey")
		return
	}

	fmt.Println("Ctrl+S hotkey registered")
	defer procUnregisterHotKey.Call(0, HOTKEY_ID)

	type MSG struct {
		Hwnd    uintptr
		Message uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
		Pt      struct{ X, Y int32 }
	}

	var msg MSG
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-a.hotkeyQuit:
			return
		case <-ticker.C:
			for {
				ret, _, _ := procPeekMessage.Call(
					uintptr(unsafe.Pointer(&msg)),
					0,
					0,
					0,
					PM_REMOVE,
				)

				if ret == 0 {
					break
				}

				if msg.Message == WM_HOTKEY && msg.WParam == HOTKEY_ID {
					go a.toggleClicking()
				}

				procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
				procDispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
			}
		}
	}
}

func (a *App) toggleClicking() {
	if a.IsRunning() {
		a.StopClicking()
	} else {
		a.StartClicking(a.lastConfig)
	}
}
