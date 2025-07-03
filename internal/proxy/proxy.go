package proxy

import (
	"io"
	"log/slog"
	"net"
)

func Run(listenAddr, backendAddr string) error {
	// create a new TCP server
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	defer ln.Close()
	for {
		// accept a new connection
		conn, err := ln.Accept()
		if err != nil {
			slog.Error("failed to accept connection", "error", err)
		}
		go handle(conn, backendAddr)
	}
}

func handle(conn net.Conn, backendAddr string) {
	defer conn.Close()

	// create a new connection to the backend server
	backendConn, err := net.Dial("tcp", backendAddr)
	if err != nil {
		slog.Error("failed to connect to backend server", "error", err)
		return
	}

	defer backendConn.Close()

	// start a goroutine to copy data from the client to the backend server
	go io.Copy(backendConn, conn)
	// copy data from the backend server to the client
	io.Copy(conn, backendConn)
}
