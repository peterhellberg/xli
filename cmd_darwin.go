// +build darwin

package main

import "os/exec"

func cmd(fn string) *exec.Cmd {
	return exec.Command("open", fn)
}
