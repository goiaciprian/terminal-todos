package main

import "terminal-todos/internal/daemon"


func main() {
	service := daemon.New()
	service.Start(true)
}