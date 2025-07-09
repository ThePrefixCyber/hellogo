# hellogo

Simple cross-platform Go program. Double click to run. Writes a text file to ../output relative to the directory of the compiled executable. Also makes a popup on Windows. 

## Requirements

- Go 1.21+ installed
- Building on macOS or Linux (not on Windows)

## Build Instructions

### Build for Windows (.exe)

```bash
GOOS=windows GOARCH=amd64 go build -o bin/hello-win.exe
```

### Build for Intel macOS

```bash
GOOS=darwin GOARCH=amd64 go build -o bin/hello-mac
```
