build_cmd:
	go build -ldflags "-X terminal-todos/cmd/commands.Version=dev-build" -o out/bin/cmd/todos.exe ./cmd
