package main

import (
	"fmt"
	"net"
	"log"
	"os"
	_"bufio"
	"io"
	_"io/ioutil"
	_"syscall"
	"time"
	_"bytes"
)

func main() {
	switch os.Args[1] {
	case "server":
		SocketServer()
	case "client":
		SocketClient()
	default:
		panic("bad command")
	}
}

func SocketServer() {
	maxsize := 1024*1024
	s, err := net.Listen("tcp4", ":5000")	
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	fmt.Printf("listening...\n")
	
	conn, err := s.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//rbuf := make([]byte, maxsize)
	for size:=1; size<=maxsize; size*=2 {
		iter := 1
		if size < 1024*2 {
			iter = 100000
		} else if size < 1024*32 {
			iter = 10000
		} else {
			iter = 100
		}
		rbuf := make([]byte, size)
		for i:=1; i<=iter; i++ {
			_ = recv(rbuf, conn)
		}
		fmt.Printf("%v[B]\n", size)
	}
}

func SocketClient() {
	maxsize := 1024*1024
	ip := os.Args[2]
	conn, err := net.Dial("tcp4", ip + ":5000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for size:=1; size<=maxsize; size*=2 {
		iter := 1
		if size < 1024*2 {
			iter = 100000
		} else if size < 1024*32 {
			iter = 10000
		} else {
			iter = 100
		}
		sbuf := make([]byte, size)
		sbuf[0] = 1
		start := time.Now()
		for i:=1; i<=iter; i++ {
			send(sbuf, conn)
		}
		end := time.Now()
		elasped_time := end.Sub(start).Seconds()
		th := 8*float64(size)*float64(iter)/elasped_time/1024/1024
		fmt.Printf("%v[B]\t, %v[s]\t, %v[loops]\t, %v[Mbps]\n", size, elasped_time, iter, th)
	}	
}

// http://golang.jp/pkg/net1
//func recv(rbuf []byte, sbuf[]byte, conn io.Reader) (int){
//func recv(rbuf []byte, sbuf []byte, conn net.Conn, f *os.File) (int){
func recv(rbuf []byte, conn net.Conn) (int) {
	recvbyte := 0	
	for {
		//recvbyte = 0	
		_, err := io.ReadFull(conn, rbuf)
		//n, err := conn.Read(rbuf)
		//_, _ = f.Write(rbuf)
		//_, err = io.Copy(ioutil.Discard, conn)
		//recvbyte += n
		//fmt.Println(recvbyte)
		break
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
	}
	return recvbyte
}

//func send(sbuf []byte, rbuf []byte, conn io.Writer) (int){
//func send(sbuf []byte, rbuf []byte, conn net.Conn) (int){
func send(sbuf []byte, conn net.Conn) {
	_, err := conn.Write(sbuf)
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}
}
