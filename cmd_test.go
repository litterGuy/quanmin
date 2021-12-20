package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"testing"
	"time"
)

func TestCMD(t *testing.T) {
	// Start a long-running process, capture stdout and stderr
	findCmd := cmd.NewCmd("java", "-version")
	statusChan := findCmd.Start() // non-blocking

	ticker := time.NewTicker(2 * time.Second)

	// Print last line of stdout every 2s
	go func() {
		for range ticker.C {
			status := findCmd.Status()
			n := len(status.Stdout)
			fmt.Println(status.Stdout[n-1])
		}
	}()

	// Stop command after 1 hour
	go func() {
		<-time.After(1 * time.Hour)
		findCmd.Stop()
	}()

	// Block waiting for command to exit, be stopped, or be killed
	finalStatus := <-statusChan
	for _,s := range finalStatus.Stdout{
		println("out is : "+s)
	}
	for _,s := range finalStatus.Stderr{
		println(s)
	}

}
