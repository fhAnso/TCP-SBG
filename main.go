package main

import (
	"WhiteRabbit/lib"
	"fmt"
	"time"
)

func EntryPoint(args *lib.Args) {
	target := fmt.Sprintf("%s:%d", args.TargetHost, args.TargetPort)
	session := lib.TcpStreamOpen(target, time.Second*5)
	banner := lib.StreamReader(session, target)
	if banner == "" {
		fmt.Println("[-] Failed to receive any data..")
	}
	fmt.Printf("[+] Response: %s", banner)
	lib.TcpStreamClose(session)
}

func main() {
	parser := lib.ArgParser()
	EntryPoint(&parser)
}
