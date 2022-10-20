//go:build windows

package constants

import "syscall"

var (
	// SYSTEMSHELL is the default shell to use on the system
	SYSTEMSHELL = "powershell"
	// SYSPROCATTR is the default process attributes to use
	SYSPROCATTR = &syscall.SysProcAttr{HideWindow: true}
)
