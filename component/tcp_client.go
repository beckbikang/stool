package component

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func TcpClientDail(ip string, port int, msg string) {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Second*5)
	if err != nil {
		log.Println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Println("Write data failed:", err.Error())
		os.Exit(1)
	}

	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		log.Println("Read data failed:", err.Error())
		os.Exit(1)
	}

	log.Println("Received message:", string(received))
	conn.Close()
}
