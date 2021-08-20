package exfiltrate

import (
	"os"
)

//function use to write data register from bash/shell session in a file
func WriteFile(data string, path string) {
	f, _ := os.Create(path)

	defer f.Close()

	f.WriteString(data)
}
