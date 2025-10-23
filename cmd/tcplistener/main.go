package main

import (
	"fmt"
	"log"
	"net"

	"silvers.rayleigh.dk/internal/request"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		for {
			r, err := request.RequestFromReader(conn)
			if err != nil {
				log.Fatal("error", "error", err)
			}
			fmt.Printf("Request line:%s\n", r.RequestLine)
			fmt.Printf("- Method:%s\n", r.RequestLine.Method)
			fmt.Printf("- Target:%s\n", r.RequestLine.RequestTarget)
			fmt.Printf("- Version:%s\n", r.RequestLine.HttpVersion)
			fmt.Printf("Headers:\n")
			r.Headers.ForEach(func(n, v string) {
				fmt.Printf("- %s: %s\n", n, v)

			})
		}

	}

}
