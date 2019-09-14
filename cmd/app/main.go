package main

import (
	"github.com/JamesHovious/w32"
	"github.com/dantheman213/quick-screen-capture/pkg/hotkeys"
	"github.com/dantheman213/quick-screen-capture/pkg/screen"
)

func main() {
	_ = hotkeys.Register()
	_ = waitForInputLoop()
}

func waitForInputLoop() (err error) {
	var msg w32.MSG
	quit := false

	for !quit {
		_ = w32.GetMessage(&msg, 0,0,0)

		switch msg.Message {
		case w32.WM_HOTKEY:
			if msg.WParam == 1 {
				screen.InitCaptureWindow()
			}
		}
	}

	return nil
}