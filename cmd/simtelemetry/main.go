package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/swerveaux/simtelemetry/internal/webserver"

	"github.com/swerveaux/simtelemetry/internal/server"
)

const port = 4843

func main() {
	fmt.Println("Starting server")
	ch := make(chan []byte)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		err := server.Run(port, ch)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		wg.Done()
	}()

	go func() {
		err := webserver.Run(10001, ch)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		wg.Done()
	}()

	wg.Wait()
	close(ch)
}
