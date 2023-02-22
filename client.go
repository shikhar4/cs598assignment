package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func write(key, value string) int {
	writeRequest := map[string]string{
		"type":  "write",
		"key":   key,
		"value": value,
	}
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	conn, err := net.DialTCP(TYPE, nil, tcpServer)

	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	data, _ := json.Marshal(writeRequest)

	_, err = conn.Write(data)
	if err != nil {
		println("Write data in write() failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("Received message:", string(received))

	conn.Close()
	return 1
}

func read(key string) int {
	writeRequest := map[string]string{
		"type": "read",
		"key":  key,
	}
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	conn, err := net.DialTCP(TYPE, nil, tcpServer)

	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	data, _ := json.Marshal(writeRequest)

	_, err = conn.Write(data)
	if err != nil {
		println("Write data in read() failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("Received message:", string(received))

	conn.Close()
	return 1
}

func main() {
	for {
		var userRequestType string
		fmt.Println("Please enter operation (write/read): ")
		fmt.Scanln(&userRequestType)

		if userRequestType == "write" {
			var key string
			var value string

			fmt.Println("Please enter key: ")
			fmt.Scanln(&key)

			fmt.Println("Please enter value: ")
			fmt.Scanln(&value)

			fmt.Println("Sending write request to servers")
			write(key, value)
		}

		if userRequestType == "read" {
			var key string

			fmt.Println("Please enter key you want to read: ")
			fmt.Scanln(&key)

			fmt.Println("Sending read request to servers")
			read(key)
		}
	}

}
