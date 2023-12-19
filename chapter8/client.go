package main

import (
	"io"
	"log"
	"net"
	"os"
	"fmt"
	"time"
)

type con struct {
	file *os.File
	con net.Conn
}

func main() {

	conn1, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn1.Close()
	conn2, err := net.Dial("tcp", "localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn2.Close()
	// f1, err := os.Create("file1.txt")
	// if err != nil {
		// log.Fatal(err)
	// }
	// defer f1.Close()
	// f2, err := os.Create("file2.txt")
	// if err != nil {
		// log.Fatal(err)
	// }
	// defer f1.Close()

	go mustCopy(os.Stdout, conn1)
	go mustCopy(os.Stdout, conn2)
	time.Sleep(10 * time.Second)
	// go mustCopy(f1, conn1)
	// mustCopy(f2, conn2)

	// c := []con{
		// {f1, conn1},
		// {f2, conn2},
	// }
// 
	// for _, val := range c {
		// go mustCopy(val.file, val.con)
	// }
}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println(src)
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
