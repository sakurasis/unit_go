package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("client listen failed, err:", err)
		return
	}
	defer func(dial net.Conn) {
		err := dial.Close()
		if err != nil {
			println("client close connection failed, err:", err)
		}
	}(dial)
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			println("")
			return
		}
		info := strings.Trim(readString, "\r\n")
		if strings.ToUpper(info) == "Q" {
			return
		}
		_, err = dial.Write([]byte(info))
		if err != nil {
			return
		}
		buf := [5120]byte{}
		i, err := dial.Read(buf[:])
		if err != nil {
			println("client receive failed, err:", err)
			return
		}
		fmt.Println(string(buf[:i]))
	}
}
