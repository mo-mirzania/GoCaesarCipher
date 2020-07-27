package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		log.Panic()
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Fprintf(conn, cesar(scanner.Text()))
	}
}

func cesar(s string) string {
	byteSlice := []byte(s)
	for i := range byteSlice {
		if byteSlice[i] <= 109 {
			byteSlice[i] += 13
		} else {
			byteSlice[i] -= 13
		}
	}
	return string(byteSlice)
}
