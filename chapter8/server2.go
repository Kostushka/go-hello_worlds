package main

import (
	"log"
	"net"
	"time"
	"flag"
	"strconv"
	"fmt"
	"strings"
	"bufio"
)

func main() {
	port := flag.Int("port", 8000, "port for server")
	flag.Parse()
	ip := "localhost:" + strconv.Itoa(*port)
	fmt.Println("ip:", ip)
	listener, err := net.Listen("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Например, обрыв соединения
			continue
		}
		go handleConn(conn) // Обработка подключения
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
