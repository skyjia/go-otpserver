package main

import (
	"./otpserver"
)

func main() {
	otpserver.StartServer(otpserver.DefaultConfig)
}
