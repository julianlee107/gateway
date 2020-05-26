package main

import (
	"fmt"
	"github.com/julianlee107/gateway/base/unpack/unpack"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Printf(" connect failed,err : %v \n", err)
		return
	}
	defer conn.Close()
	err = unpack.Encode(conn, "hello world")
	if err != nil {
		fmt.Printf(" encode failed,err : %v \n", err)
		return
	}
}
