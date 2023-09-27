package main

import (
	"fmt"
	"github.com/ellypaws/stencil/gui" // <-- Add this import
	"os"
	"os/signal"
	"syscall"
)

func handleInterrupts() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\nReceived Ctrl+C, exiting...")
		os.Exit(0)
	}()
}

func main() {
	handleInterrupts() // Handles Ctrl + C
	gui.NewApp()       // ---> Add this line at the bottom
}
