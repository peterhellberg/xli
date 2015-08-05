// +build windows

package main

import "os/exec"

func cmd(fn string) *exec.Cmd {
	return exec.Command("rundll32", "C:\\\\Windows\\\\System32\\\\shimgvw.dll,imageview_fullscreen", fn)
}
