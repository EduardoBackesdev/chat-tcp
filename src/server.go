package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Server() {
	ch := make(chan net.Conn, 5)
	fmt.Println("Servidor criado!")
	go ListenConnection(ch)
	fmt.Println("Conexão aceita...")
	
	for {
		conexao := <-ch
		mes, err := bufio.NewReader(conexao).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		fmt.Print("CHAT GLOBAL: ",mes)
	}

}
