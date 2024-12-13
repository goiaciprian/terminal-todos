package daemon

import (
	"runtime"

	"terminal-todos/internal"
	"terminal-todos/internal/daemon/windows"
)

type DaemonService struct {
	OS string
}

func (s *DaemonService) Start(debug bool) {
	switch s.OS {
		case "windows":
			windows.RunService(internal.APP_NAME, debug)
	}
}

func New() DaemonService {
	return DaemonService{
		OS: runtime.GOOS,
	}
}
