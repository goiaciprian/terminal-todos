package notification

import (
	"fmt"
	"os"
	"runtime"

	"terminal-todos/internal"
	"terminal-todos/internal/config"

	"gopkg.in/toast.v1"
)

func windowsReminder(m string) {
	notification := toast.Notification{
		AppID: internal.DISPLAY_NAME,
		Message: m,
		Icon: config.Instance.ICON_PATH,
	}

	err := notification.Push()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func SendReminder (message string) {
	switch runtime.GOOS {
		case "windows": {
			windowsReminder(message)
		}
	}
}
