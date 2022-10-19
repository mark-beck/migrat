package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func readPipe(pipe io.ReadCloser, out chan string) {
	reader := bufio.NewReader(pipe)
	for {

		output, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading shell: ", err.Error())
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
			log.Println("Error writing to shell: ", err.Error())
		}
	}
}
