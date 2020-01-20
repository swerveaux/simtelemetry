package server

import (
	"fmt"
	"net"
	"strconv"
)

// Run takes a port number and starts listening on that port.
func Run(port int, ch chan []byte) error {
	pc, err := net.ListenPacket("udp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	defer pc.Close()

	buffer := make([]byte, 1024)
	for {
		numBytes, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			fmt.Printf("Got error: %v\n", err)
			return err
		}
		fmt.Printf("Received %d bytes from %v: %v\n", numBytes, addr, buffer[:numBytes])
		ch <- buffer[:numBytes]
	}

	return nil
}
