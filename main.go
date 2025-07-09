package main

import (
    "os"
    "os/exec"
    "path/filepath"
)

func main() {
    exePath, err := os.Executable()
    if err != nil {
        return
    }
    baseDir := filepath.Dir(exePath)
    outputPath := filepath.Join(baseDir, "..", "output", "hello-output.txt")

    os.MkdirAll(filepath.Dir(outputPath), 0755)

    f, err := os.Create(outputPath)
    if err != nil {
        return
    }
    defer f.Close()

    msg := "Hello from Go!"
    f.WriteString(msg)

    // Show a popup dialog (Windows only)
    exec.Command("powershell", "-Command",
        `Add-Type -AssemblyName PresentationFramework;`+
            `[System.Windows.MessageBox]::Show("`+msg+`")`,
    ).Run()
}
