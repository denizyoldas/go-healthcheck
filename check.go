package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("DOWN: %s", err.Error())
	} else {
		defer conn.Close()
		status = fmt.Sprintf("UP")
	}

	return status
}