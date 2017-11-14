package main

import (
	"os/exec"
	"log"
	"flag"
	"os"
	"fmt"
	"strings"
	"time"
	"bytes"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage\n  %s [Options] Arguments...\nOptions\n", os.Args[0])
		flag.PrintDefaults()
	}
	var (
		display = flag.String("display", ":0", "DISPLAY environment variable")
	)
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	environment := []string{fmt.Sprintf("DISPLAY='%s'", *display)}
	args := append(environment, flag.Args()...)
	cmd := exec.Command("bash", "-c", strings.Join(args, " "))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.Println("Executing...", cmd.Args)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	log.Println("Running at PID", cmd.Process.Pid)
	time.Sleep(100 * time.Millisecond)

	if output := stdout.String(); output != "" {
		log.Println(output)
	}
	if output := stderr.String(); output != "" {
		log.Println(output)
	}

	cmd.Process.Release()
}
