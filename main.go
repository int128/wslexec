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
	"bufio"
)

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
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Panic(err)
	}

	if err := cmd.Start(); err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(translateWslPath(scanner.Text()))
	}

	if err := cmd.Wait(); err != nil {
		status := 1
		if exitError, ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
				status = waitStatus.ExitStatus()
			}
		}
		os.Exit(status)
	}
}

var wslDrivePathPattern = regexp.MustCompile("/mnt/([[:alpha:]])/")

func translateWslPath(line string) string {
	return wslDrivePathPattern.ReplaceAllString(line, "$1:/")
}

func translateWindowsPathInArgs(windowsPathArgs []string) []string {
	unixPathArgs := make([]string, len(windowsPathArgs))
	for i, windowsPathArg := range windowsPathArgs {
		unixPathArgs[i] = translateWindowsPathInArg(windowsPathArg)
	}
	return unixPathArgs
}

var windowsDrivePathPattern = regexp.MustCompile("([[:alpha:]]):\\\\")

func translateWindowsPathInArg(arg string) string {
	if windowsDrivePathPattern.FindStringIndex(arg) != nil {
		driveReplaced := windowsDrivePathPattern.ReplaceAllStringFunc(arg, func (drivePath string) string {
			m := windowsDrivePathPattern.FindStringSubmatch(drivePath)
			drive := strings.ToLower(m[1])
			return fmt.Sprintf("/mnt/%s/", drive)
		})
		backslashReplaced := strings.Replace(driveReplaced, "\\", "/", -1)
		return backslashReplaced
	} else {
		return arg
	}
}
