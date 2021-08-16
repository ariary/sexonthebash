package main

import (
	"fmt"
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

func readStdin() {

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	for {
		if fi.Size() > 0 {
			fmt.Println("there is something to read")
		}
	}
}

func main() {
	fmt.Println("before bash")
	go readStdin()
	bash := exec.Command("/bin/bash")
	bash.Stderr = os.Stderr
	bash.Stdout = os.Stdout
	bash.Stdin = os.Stdin

	bash.Run()
	fmt.Println("after bash")
}
