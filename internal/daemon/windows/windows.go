package windows

import (
	"fmt"

	"terminal-todos/internal/notification"
	"terminal-todos/internal/todo"

	"github.com/krayzpipes/cronticker/cronticker"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

type service struct{}

var wlog debug.Log

func (m *service) Execute(args []string, r <-chan svc.ChangeRequest, status chan<- svc.Status) (bool, uint32) {

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue

	tick, err := cronticker.NewTicker("15/30 */1 * * *")
	if err != nil{
		wlog.Error(1, fmt.Sprintf("error ticker #%s", err.Error()))
		return false, 1
	}

	status <- svc.Status{State: svc.StartPending}

	status <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

loop:
	for {
		select {
		case <-tick.C:
			todosCount, err := todo.Count()
			if err != nil {
				wlog.Error(1, err.Error())
				wlog.Error(1, "[cron] error reading todos")
			}
			wlog.Info(1, fmt.Sprintf("[cron] %d found todos", todosCount))
			if todosCount != 0 {
				notification.SendReminder(fmt.Sprintf("You have %d todos uncompleted", todosCount))
			}
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				status <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				break loop
			case svc.Pause:
				status <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
			case svc.Continue:
				status <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
			default:
				wlog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}

	status <- svc.Status{State: svc.StopPending}
	return false, 0
}

func RunService(name string, isDebug bool) {
	if isDebug {
		wlog = debug.New(name)
	} else {
		err := eventlog.InstallAsEventCreate(name, eventlog.Info|eventlog.Warning|eventlog.Error)
		if err != nil {
			fmt.Printf("error installing as event, %s", err.Error())
		}

		wlog, err = eventlog.Open(name)
		if err != nil {
			fmt.Printf("error oppening event, %s", err.Error())
		}
	}

	if isDebug {
		err := debug.Run(name, &service{})
		if err != nil {
			wlog.Error(1, fmt.Sprintf("Error running service in debug mode. \n %s", err.Error()))
		}
	} else {
		err := svc.Run(name, &service{})
		if err != nil {
			wlog.Error(1, fmt.Sprintf("Error running service in Service Control mode. \n %s", err.Error()))
		}
	}
}
