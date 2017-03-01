package main

//网络时钟服务
import (
	"net"
	"log"
	"io"
	"time"
)

const host = "localhost:12345"

func main() {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05 PM Mon\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
