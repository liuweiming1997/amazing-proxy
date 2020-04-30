package main

import (
	// "encoding/binary"
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	ServerAddress = "0.0.0.0"
	ServerPort = "9090"
)

func init() {
	log.Info("Running amazing-proxy with pid = ", os.Getpid())
	log.Info(fmt.Sprintf("Process listen host = %s:%s", ServerAddress, ServerPort))
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", ServerAddress, ServerPort))
	if err != nil {
		log.Error("Can not resolve host, exit...", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Error("Error listening:",  err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 4)
	dataSize, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Error("Read data from UDP fail", err)
		return
	}
	log.Info(dataSize, remoteAddr, data)
	conn.WriteToUDP(data, remoteAddr)
}
