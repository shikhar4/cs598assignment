package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST      = "localhost"
	PORT      = "8080"
	TYPE      = "tcp"
	TIMESTAMP = 0
)

var store = make(map[string]string)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// incoming request
	request := map[string]string{}
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(buffer[:n], &request)

	if request["type"] == "write" {
		fmt.Println("Write Request")
		key := request["key"]
		value := request["value"]
		store[key] = value

		responseStr := fmt.Sprintf("Wrote value %v into key %v. Received time: %v", value, key, time.Now().Format(time.ANSIC))
		fmt.Println(store)
		conn.Write([]byte(responseStr))
	}

	if request["type"] == "read" {
		fmt.Println("Read Request")
		key := request["key"]
		value := store[key]
		responseStr := fmt.Sprintf("%v", value)
		conn.Write([]byte(responseStr))
	}
	// close conn
	conn.Close()
}
