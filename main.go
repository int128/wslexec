package main

import (
	"os/exec"
	"log"
	"os"
	"strings"
	"path/filepath"
	"syscall"
	"regexp"
	"fmt"
)

var windowsDrivePathPattern = regexp.MustCompile("(\\w):\\\\")

func main() {
	baseName := filepath.Base(os.Args[0])
	if ! strings.HasPrefix(baseName, "wsl") {
		log.Println("Basename does not have prefix: ", baseName)
		log.Fatal("Rename this binary to command name with prefix `wsl`. For example, rename to `wslgit` to run git command on WSL.")
	}

	commandName := strings.TrimSuffix(strings.TrimPrefix(baseName, "wsl"), ".exe")
	commandArgs := translateWindowsPathInArgs(os.Args[1:])
	commandLine := append([]string{commandName}, commandArgs...)

	cmd := exec.Command("wsl", commandLine...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		status := 1
		if exitError, ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
				status = waitStatus.ExitStatus()
			}
		}
		os.Exit(status)
	}
}

func translateWindowsPathInArgs(windowsPathArgs []string) []string {
	unixPathArgs := make([]string, len(windowsPathArgs))
	for i, windowsPathArg := range windowsPathArgs {
		unixPathArgs[i] = translateWindowsPathInArg(windowsPathArg)
	}
	return unixPathArgs
}

func translateWindowsPathInArg(arg string) string {
	driveReplaced := windowsDrivePathPattern.ReplaceAllStringFunc(arg, func (drivePath string) string {
		m := windowsDrivePathPattern.FindStringSubmatch(drivePath)
		drive := strings.ToLower(m[1])
		return fmt.Sprintf("/mnt/%s/", drive)
	})
	backslashReplaced := strings.Replace(driveReplaced, "\\", "/", -1)
	return backslashReplaced
}
