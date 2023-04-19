package readFile

import (
	"io"
	"os"
)


func PrintToTerminal(r io.Reader){
	io.Copy(os.Stdout , r)
}