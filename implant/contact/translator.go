package contact

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

func ReadMessage(conn net.Conn, key []byte) (proto.Message, error) {
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
		message = &Ident{}
	case 1:
		message = &HeartbeatResponse{}
	case 2:
		message = &Error{}
	case 10:
		message = &ModuleStart{}
	case 11:
		message = &ModuleInput{}
	case 12:
		message = &ModuleOutput{}
	case 13:
		message = &ModuleList{}
	case 20:
		message = &TakeScreenshot{}
	case 21:
		message = &Screenshot{}
	default:
		return nil, &cantDecodeMessageError{}

	}
	if err := proto.Unmarshal(bytes, message); err != nil {
		return nil, &cantDecodeMessageError{}
	}
	return message, nil

}

func SendMessage(conn net.Conn, message proto.Message, key []byte) error {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	var type_byte byte

	switch message.(type) {
	case *Ident:
		type_byte = 0
	case *HeartbeatResponse:
		type_byte = 1
	case *Error:
		type_byte = 2
	case *ModuleStart:
		type_byte = 10
	case *ModuleInput:
		type_byte = 11
	case *ModuleOutput:
		type_byte = 12
	case *ModuleList:
		type_byte = 13
	case *TakeScreenshot:
		type_byte = 20
	case *Screenshot:
		type_byte = 21
	}

	bytes = append(bytes, type_byte)
	bytes = encrypt(bytes, key)

	return sendMessageBytes(conn, bytes)
}

type cantDecodeMessageError struct{}

func (e *cantDecodeMessageError) Error() string {
	return "cant Decode Message"
}
