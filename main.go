package main

import (
	"os/exec"
	"log"
	"flag"
	"os"
	"fmt"
	"strings"
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

	cmd := exec.Command("bash", "-c", strings.Join(flag.Args(), " "))
	cmd.Env = append(os.Environ(), fmt.Sprintf("DISPLAY=%s", display))

	log.Println("Executing...", cmd.Args)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Running... PID", cmd.Process.Pid)
}
