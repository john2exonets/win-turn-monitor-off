//go:build windows
// +build windows

package main

//
//  turnOffMonitor  --  Shuts down the signal to the monitor on Windows.  A better power saver
//    over using 'scrnsave.scr /s' to just blank the screen.
//
//  John D. Allen
//  September, 2022
//

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"golang.org/x/sys/windows"
)

func main() {
	user32DLL := windows.NewLazyDLL("user32.dll")
	procPostMsg := user32DLL.NewProc("PostMessageW")

	_, _, err := procPostMsg.Call(0xFFFF, 0x0112, 0xF170, 2)

	//
	// I get a "Access is denied." error when I run this from a user that does
	// not have SYSTEM rights...but it still works anyway for some reason.
	if err != nil && !strings.Contains(err.Error(), "Access is denied.") {
		log.Errorf("Error returned from PostMsg(): %v", err)
	}
}
