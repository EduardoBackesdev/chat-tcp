package src

import (
	"fmt"
	"net"
	"os"
)

func Connection() {
	// red := color.New(color.FgRed).Add(color.Underline).SprintFunc()
	// green := color.New(color.FgGreen).Add(color.Underline).SprintFunc()
	// blue := color.New(color.FgBlue).Add(color.Underline).SprintFunc()
	// yellow := color.New(color.FgYellow).Add(color.Underline).SprintFunc()
	// white := color.New(color.FgWhite).Add(color.Underline).SprintFunc()
	// cores := []interface{}{
	// 	red,green,blue,yellow,white,
	// }
	_, err := net.Dial("tcp", "192.168.0.105:8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	// nome := ""
	// if nome == "" {
	// 	nome = User()
	// }
	// ran := cores[rand.IntN(len(cores))]
	// for {
	// 	leitor := bufio.NewReader(os.Stdin)
	// 	texto, textoErr := leitor.ReadString('\n')
	// 	if textoErr != nil {
	// 		fmt.Println(textoErr)
	// 		os.Exit(3)
	// 	}

	// 	fmt.Fprintf(conexao, nome+": "+texto+"\n")
	// 	fmt.Println("")
	// }
}
