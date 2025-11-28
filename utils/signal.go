package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HandleCtrlC() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\nProgram interrupted. Exiting cleanly...")
		os.Exit(0)
	}()
}
