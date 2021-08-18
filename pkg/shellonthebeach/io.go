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

func WriteCharAtIndex(str string, c rune, index int) (new string, err error) {
	sz := len(str)
	if index == sz {
		new = str + string(c)

	} else if index < sz {
		new = str[:index] + string(c) + Dir[index:]
	}

	return new, fmt.Errorf("WriteCharAtIndex: Index out of range")
}

func GetCommandAndListenArrow() (cmd string) {
	//keyboard listener
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	//For interativeness with arrow
	index := 0 //writing position in cmd

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
			index += 1
		case keyboard.KeyBackspace:
			sz := len(cmd)
			if sz > 2 {
				cmd = cmd[:sz-1]
				fmt.Printf("\r$ %s ", cmd) //does not handle multiple-line case)
				index -= 1
			}
		// case keyboard.KeyArrowUp:
		// case keyboard.KeyArrowDown:
		case keyboard.KeyArrowLeft:
			index -= 1
			fmt.Printf("\r$ %s", cmd)
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
