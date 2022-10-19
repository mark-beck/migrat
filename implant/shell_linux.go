//go:build linux

package main

import (
	"log"
	"os/exec"
	"syscall"
)

type ShellError struct{}

func (e *ShellError) Error() string {
	return "shell error"
}

func init_shell() (*Module, error) {
	log.Println("loading system shell")
	command := exec.Command("bash")
	command.SysProcAttr = &syscall.SysProcAttr{}
	outpipe, err1 := command.StdoutPipe()
	inpipe, err2 := command.StdinPipe()
	errpipe, err3 := command.StderrPipe()

	if err1 != nil || err2 != nil || err3 != nil {
		log.Println(err1.Error(), err2.Error(), err3.Error())
		return nil, &ShellError{}
	}

	module := &Module{
		finished: make(chan bool, 1),
		name:     "shell",
		external: true,
		args:     []string{},
		in:       make(chan string, 1),
		out:      make(chan string, 1),
		err:      make(chan string, 1),
	}

	go readPipe(outpipe, module.out)
	go readPipe(errpipe, module.err)
	go writePipe(inpipe, module.in)
	go func() {
		command.Run()
		module.finished <- true
		log.Println("module finished: ", module.name)
	}()
	log.Println("shell started")

	return module, nil
}
