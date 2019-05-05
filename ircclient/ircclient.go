package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Redes-2019/connection"
	"github.com/Redes-2019/userinterface"
)

func main() {

	// Pega informacões de usuário
	fmt.Println("Iniciando Cliente.")
	fmt.Println("Identificando usuário")
	user, conn := userinterface.ReadUserData()
	channel := "> "
	
	// Entra em loop para ler os comandos do usuário
	for {
		command := userinterface.ReadCommand(channel)
		if command == nil {
			fmt.Println("Erro: comando invalido")
			} else {
				validStructure := userinterface.VerifyStructure(command)
				if validStructure == false {
					fmt.Println("Erro: Estrutura de comando incorreta")
				} else {
					connection.HandleConnection(command)
				}
			}
	}

	// Conecta ao servidor
	connSocket := connection.OpenSocket(conn)

	// Usa o socket para conectar o client
	client := connection.NewClient(connSocket, user, conn)

	go client.Receive()
	client.Connect(user)

	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		strings.TrimRight(message, "\n")
		message += "\r\n"
		fmt.Print("SENT:", message)
		connSocket.Write([]byte(message))
	}

	fmt.Println("Client disconnected. Terminating.")
}
