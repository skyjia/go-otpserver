package main

import (
	"github.com/skyjia/go-otpserver/otpserver"
)

func main() {
	otpserver.StartServer(otpserver.DefaultConfig)
}
