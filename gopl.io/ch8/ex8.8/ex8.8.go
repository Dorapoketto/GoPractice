package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	var message = make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(10 * time.Second):
			c.Close()
		case msg := <-message:
			wg.Add(1)
			go func(c net.Conn, shout string, delay time.Duration) {
				defer wg.Done()
				fmt.Fprintln(c, "\t", strings.ToUpper(shout))
				time.Sleep(delay)
				fmt.Fprintln(c, "\t", shout)
				time.Sleep(delay)
				fmt.Fprintln(c, "\t", strings.ToLower(shout))
			}(c, msg, 1*time.Second)
		}
	}()

	for input.Scan() {
		text := input.Text()
		message <- text
	}
	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	//cw := c.(*net.TCPConn)
	//cw.CloseWrite()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
