package main

import "os"

func main() {
	if len(os.Args) > 1 {
		cmd(os.Args[1]).Run()
	}
}
