package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kinvolk/procconnector"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Collect new processes for a determined amount of time.\n\tUsage: %s NSEC\n", os.Args[0])
		os.Exit(1)
	}

	waitTime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid argument %q: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	pc, err := procconnector.New(false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n")
		os.Exit(1)
	}

	if !pc.IsRunning() {
		fmt.Fprintf(os.Stderr, "proc connector is not running\n")
		os.Exit(1)
	}

	time.Sleep(time.Duration(waitTime) * time.Second)
	pc.Walk(func(pid procconnector.Process) {
		fmt.Println(pid)
	})
}
