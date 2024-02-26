package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		panic("cannot start the server")
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		// for handling concurrent clients
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	for {
		r := bufio.NewReader(conn)
		resp := NewResp(r)
		val, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		if val.typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(val.array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(val.array[0].bulk)
		args := val.array[1:]

		writer := NewWriter(conn)
		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{
				typ: "string",
				str: "",
			})
		}
		result := handler(args)
		writer.Write(result)
	}
}
