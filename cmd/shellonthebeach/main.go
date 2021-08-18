package main

import (
	"fmt"
	"sexonthebash/pkg/command"
	"sexonthebash/pkg/shellonthebeach"
)

// see:
//https://stackoverflow.com/questions/43965556/how-to-read-a-key-in-go-but-continue-application-if-no-key-pressed-within-x-seco

//Error handling for debug purpose. To be stealthed withdraw them
func main() {
	for {
		shellonthebeach.Prefix()
		cmd := shellonthebeach.GetCommand()
		command.ManageSpecialCase(cmd)
		output := command.Exec(cmd)
		fmt.Println(output)
	}
}
