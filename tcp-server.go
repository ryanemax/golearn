// 182.92.5.210:4567
package main

import (
	"bufio"
	"fmt"
	"net"
)

func Echo(c net.Conn) {

	defer c.Close()
	for {
		line, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Printf("Failure to read:%s\n", err.Error())
			return
		}
		fmt.Println(line)
		_, err = c.Write([]byte(line))
		if err != nil {
			fmt.Printf("Failure to write: %s\n", err.Error())
			return
		}
	}

}

func main() {
	fmt.Println("Server is ready...\n")
	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println(err.Error, "\n")
	}
	for {
		if c, err := ln.Accept(); err == nil {
			go Echo(c) //new thread
		}
	}
}
