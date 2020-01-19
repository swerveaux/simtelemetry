package main

import (
	"fmt"

	"github.com/swerveaux/simtelemetry/internal/server"
)

const port = 8080

func main() {
	fmt.Println("Starting server")
	server.Run(port)
}

