package src

import (
	"fmt"
	"net"
	"os"
)

func ListenConnection() {
	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	// aceitando conexões
	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close()
}