package module

import (
	"log"
	"migrat/constants"
	"os/exec"
)

type ProcessError struct{}

func (e *ProcessError) Error() string {
	return "Process error"
}

func process_start(name string, processname string, args []string) (*Module, error) {
	log.Println("process_start --> name: ", name, " processname: ", processname, " args: ", args)
	command := exec.Command(processname, args...)
	command.SysProcAttr = constants.SYSPROCATTR
	outpipe, err1 := command.StdoutPipe()
	inpipe, err2 := command.StdinPipe()
	errpipe, err3 := command.StderrPipe()

	if err1 != nil || err2 != nil || err3 != nil {
		log.Println(err1.Error(), err2.Error(), err3.Error())
		return nil, &ProcessError{}
	}

	module := &Module{
		finished: make(chan bool, 1),
		name:     name,
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
	log.Println("process started")

	return module, nil
}
