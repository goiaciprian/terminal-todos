# Terminal Todos 
[![Release](https://github.com/goiaciprian/terminal-todos/actions/workflows/release.yml/badge.svg)](https://github.com/goiaciprian/terminal-todos/actions/workflows/release.yml)

This is a small todo project for terminal

## Instalation
1. Download executable
2. Add exe to path
3. Run `todos install`
4. Optional - start service
   1. windows - in an admin powershell run `sc create terminaltodos binPath= "<PATH>\todos.exe /serve" DisplayName= 'Terminal Todos' start= auto` then `sc start terminaltodos`
