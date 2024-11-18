build_cmd:
	go build -ldflags "-X terminal-todos/cmd/commands.Version=dev-build" -o out/bin/cmd/todos.exe ./cmd

build_cmd_debug:
	go build -gcflags=all="-N -l" -ldflags "-X terminal-todos/cmd/commands.Version=dev-build" -o out/bin/cmd/todos_debug.exe ./cmd