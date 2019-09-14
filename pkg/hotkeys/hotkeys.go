package hotkeys

import (
    "fmt"
    "github.com/JamesHovious/w32"
)

// https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
func Register() error {
    err := w32.RegisterHotKey(0, 1, w32.MOD_ALT | w32.MOD_SHIFT | w32.MOD_NOREPEAT, 0x34) // bind SHIFT+ALT+4
    if err != nil {
        fmt.Println("Error registering hotkey:", err.Error())
        return err
    }

    return nil
}
