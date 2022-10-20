//go:build linux

package constants

import "syscall"

var (
	// SYSTEMSHELL is the default shell to use on the system
	SYSTEMSHELL = "bash"
	// SYSPROCATTR is the default process attributes to use
	SYSPROCATTR = &syscall.SysProcAttr{}
)
