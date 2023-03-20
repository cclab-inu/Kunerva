package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cclab-inu/Kunerva/src/core"
)

func runCommand(commandStr string) {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)

	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "cronjob":
		core.StartCronJob()
	case "onetime":
		core.StartToDiscoverNetworkPolicies()
	default:
		println("cronjob or onetime?")
	}
}

func main() {
	time.Sleep(5 * time.Millisecond)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		runCommand(cmdString)
	}
}
