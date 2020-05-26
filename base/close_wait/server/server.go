package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func(conn net.Conn) {
			// 没有正确关闭会引起，close wait
			defer conn.Close()
			for {
				var buf [128]byte
				n, err := conn.Read(buf[:])
				if err != nil {
					fmt.Printf("read from connect buff err: %v \n", err)
					break
				}
				str := string(buf[:n])
				fmt.Printf("receive from client ,data: %v \n", str)
			}
		}(conn)
	}
}
