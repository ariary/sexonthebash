package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
)

// func test() (in string, out string, err error) {
// 	// Create arbitrary command.

// 	return mw
// }

func main() {
	fmt.Println("before bash")
	c := exec.Command("bash")

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// // Set stdin in raw mode.
	// oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	// NOTE: The goroutine will keep reading until the next keystroke before returning.

	var outBuffer bytes.Buffer
	// var errBuffer bytes.Buffer
	var inBuffer bytes.Buffer

	mwOut := io.MultiWriter(os.Stdout, &outBuffer)
	in := io.TeeReader(os.Stdin, &inBuffer)
	go func() { _, _ = io.Copy(ptmx, in) }()
	_, _ = io.Copy(mwOut, ptmx)

	fmt.Println("after bash")
	fmt.Println("Captured output", outBuffer.String())
	// fmt.Println("Captured error", errBuffer.String())
	fmt.Println("Captured input", inBuffer.String())
}
