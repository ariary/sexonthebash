package shellonthebeach

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

const prefix = "$ "

//Get command from user input
func GetCommand() (cmd string) {
	for {
		buf := bufio.NewReader(os.Stdin)
		c, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		cmd += strings.Trim(c, "\n")
		if !strings.HasSuffix(cmd, "\\") {
			return cmd
		} else {
			//Multiple-line command
			cmd = strings.Trim(cmd, "\\")
		}
	}
}

func GetCommandAndListenArrow() (cmd string) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for { //while(1)
		event := <-keysEvents
		if event.Err != nil {
			fmt.Println(event.Err)
		}
		switch event.Key {
		case keyboard.KeyEnter:
			if !strings.HasSuffix(cmd, "\\") { //just check if cmdline does not finish with \
				fmt.Println()
				return cmd
			} else {
				//Multiple-line command
				fmt.Println()
				cmd = strings.Trim(cmd, "\\")
			}
		case keyboard.KeySpace:
			//event.Rune for space could be x00 => error
			space := " "
			fmt.Print(space)
			cmd += space
		case keyboard.KeyBackspace:
			sz := len(cmd)
			if sz > 2 {
				cmd = cmd[:sz-1]
				fmt.Printf("\r$ %s", cmd) //does not handle multiple-line case)
			}
		// case keyboard.KeyArrowUp:
		// case keyboard.KeyArrowDown:
		// case keyboard.KeyArrowLeft:
		// case keyboard.KeyArrowRight:
		default:
			fmt.Printf("%s", string(event.Rune))
			cmd += string(event.Rune)
		}
	}
}

//Print the prefix of the shell
func Prefix() {
	fmt.Print(prefix)
}
