package main

import (
	"io"
	"log"
	"net"
	"os"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		fmt.Println("go1")
		io.Copy(os.Stdout, conn) // Примечание: игнорируем ошибки
		log.Println("done")
		done <- true // Сигнал главной go-подпрограмме
	}()
	mustCopy(conn, os.Stdin)
	fmt.Println("go2")
	conn.Close()
	<-done
	// Ожидание завершения фоновой go-подпрограммы
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
