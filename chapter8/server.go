package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	"strconv"
	"fmt"
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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // Например, отключение клиента
		}
		time.Sleep(1 * time.Second)
	}
}
