package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	doSend()

}
func doSend()  {
	conn,err := net.Dial("tcp","0.0.0.0:9000")

	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	for{
		input,err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
	// 没有正确关闭时，会导致close wait
	defer conn.Close()
}