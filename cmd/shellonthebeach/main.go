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
	var historic shellonthebeach.Historic
	for {
		shellonthebeach.Prefix()
		cmd := shellonthebeach.GetCommandInteractive(&historic)
		command.ManageSpecialCase(cmd)
		output := command.Exec(cmd)
		fmt.Println(output)
	}
}
