package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		println("listen failed，err", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept data failed, err:", err)
			continue
		}
		go handler(conn)
	}

}

/**
handler the function deal some datas from the client.
*/
func handler(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("it occures some errors on closing the connect TCP that is ", err)
		}
	}(conn)
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		read, err := reader.Read(buf[:])
		if err != nil {
			println("receive from client failed, please check the data with connecting the admin.err:", err)
			break
		}
		s := string(buf[:read])
		println("server receive :", s)
		_, err = conn.Write([]byte(s))
		if err != nil {
			print("write the data fail, the err is ", err)
			return
		}
	}
}
