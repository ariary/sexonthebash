package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//go routine which read stdin permently (and print it for now);
// Maybe use mutex, safe read to avoid race condition
// Once it is ok the same thing for stdout

// see:
//https://eli.thegreenplace.net/2020/faking-stdin-and-stdout-in-go/
//StdinPipe for exec command
//https://zetcode.com/golang/pipe/
//https://coderwall.com/p/zyxyeg/golang-having-fun-with-os-stdin-and-shell-pipes
//https://stackoverflow.com/questions/50788805/how-to-read-from-stdin-with-goroutines-in-golang

func main() {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fmt.Println("before bash")

	bash := exec.Command("/bin/bash")
	bash.Stderr = os.Stderr
	bash.Stdout = os.Stdout
	bash.Stdin = os.Stdin

	bash.Run()
	fmt.Println("after bash")
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	fmt.Printf("Captured: %s", out)
}
