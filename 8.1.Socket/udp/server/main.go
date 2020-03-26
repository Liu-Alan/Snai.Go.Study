package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	len, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println(string(buf[0:len]))

	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func main() {
	service := ":7070"
	updAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", updAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}
