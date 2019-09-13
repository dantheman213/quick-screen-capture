package main

import (
	"fmt"
	"github.com/JamesHovious/w32"
)

func main() {
	_ = registerHotKeys()
	_ = waitForInputLoop()
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
func registerHotKeys() error {
	err := w32.RegisterHotKey(0, 1, w32.MOD_ALT | w32.MOD_SHIFT | w32.MOD_NOREPEAT, 0x34) // bind SHIFT+ALT+4
	if err != nil {
		fmt.Println("Error registering 3:", err.Error())
		return err
	}

	return nil
}

func waitForInputLoop() (err error) {
	var msg w32.MSG
	quit := false
	for !quit {
		_ = w32.GetMessage(&msg, 0,0,0)

		switch msg.Message {
		case w32.WM_HOTKEY:
			if msg.WParam == 1 {
				w32.MessageBox(0, "Event occurred!", "Event", 0)
			}
		}
	}

	return nil
}