package main

import (
	"os"
	"os/exec"
)

func GetEvnWithDefaultVal(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	} else {
		return defaultVal
	}
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
