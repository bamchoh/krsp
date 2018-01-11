package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "args is not enough...")
		os.Exit(1)
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid pid : %s", os.Args[1])
		os.Exit(1)
	}

	p := os.Process{Pid: pid}
	if err := p.Kill(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
