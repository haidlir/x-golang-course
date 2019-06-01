package main

import (
	"fmt"
)

const (
	ifconfig = iota // 1 
	netstat			// 2
	ping			// 3
)

func main() {
	command := ping
	switch command {
	case ifconfig:
		fmt.Println("Print interface configuration")
	case netstat:
		fmt.Println("Print network status")
	case ping:
		fmt.Println("Check connection")
	default:
		fmt.Println("Unknown command")
	}
}