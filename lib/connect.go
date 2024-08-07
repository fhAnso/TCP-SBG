package lib

import (
	"fmt"
	"net"
	"os"
	"time"
)

func TcpStreamOpen(target string, timeout time.Duration) net.Conn {
	fmt.Printf("[*] Opening TCP stream to %s..\n", target)
	stream, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		fmt.Printf("[-] %s\n", err)
		os.Exit(-1)
	}
	fmt.Println("[*] Session opened!")
	return stream
}

func TcpStreamClose(session net.Conn) {
	session.Close()
	fmt.Println("[*] Session closed")
}

func StreamReader(session net.Conn, target string) string {
	buffer := make([]byte, 1024)
	httpRequest := "GET / HTTP/1.1\r\nHost: %s\r\n\r\n"
	fmt.Fprintf(session, httpRequest, target)
	fmt.Println("[*] Reading data from stream..")
	reader, _ := session.Read(buffer)
	banner := string(buffer[:reader])
	return banner
}
