package main

import (
	"fmt"
)

func main() {
	command := "ping"
	switch command {
	case "ifconfig":
		fmt.Println("Print interface configuration")
	case "netstat":
		fmt.Println("Print network status")
	case "ping":
		fmt.Println("Check connection")
	default:
		fmt.Println("Unknown command")
	}
}
