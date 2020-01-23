package main

import (
	"fmt"
	"os"
	"sync"

	flag "github.com/spf13/pflag"
	"github.com/swerveaux/simtelemetry/internal/server"
	"github.com/swerveaux/simtelemetry/internal/webserver"
)

type Config struct {
	UDPPort int
	HTTPPort int
}

func main() {
	var c Config

	flag.IntVar(&c.HTTPPort, "http-port", 10001, "Port to run http server on for displaying telemetry")
	flag.IntVar(&c.UDPPort, "udp-port", 4843, "Port to run UDP listener on.")
	flag.Parse()
	fmt.Printf("Starting UDP server on %d, HTTP server on %d\n", c.UDPPort, c.HTTPPort)
	ch := make(chan []byte)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		err := server.Run(c.UDPPort, ch)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		wg.Done()
	}()

	go func() {
		err := webserver.Run(c.HTTPPort, ch)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		wg.Done()
	}()

	wg.Wait()
	close(ch)
}
