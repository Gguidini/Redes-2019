package main

import (
	"fmt"

	"github.com/Redes-2019/connection"
	"github.com/Redes-2019/userinterface"
)

func main() {
	fmt.Println("Iniciando Cliente.")
	// Pega informacões de usuário
	fmt.Println("Identificando usuário")
	user, conn := userinterface.ReadUserData()
	// Criando socket com o servidor
	connSocket := connection.OpenSocket(conn)
	// Criando cliente IRC
	client := connection.NewClient(connSocket, user, conn)
	// Conectando cliente ao servidor
	ok := client.Connect()
	if !ok {
		fmt.Println("Encerrando.")
	}

	go client.ListenServer()
	go listenUser(client)
	ok = true
	for ok {
		select {
		case msg, open := <-client.DataFromServer:
			if !open {
				ok = false
			} else {
				fmt.Println("RECEIVED:", msg)
			}
		case msg, _ := <-client.DataFromUser:
			client.HandleConnection(msg)
		}
	}

	fmt.Println("Client disconnected. Terminating.")
}

func listenUser(client *connection.IrcClient) {
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
				client.DataFromUser <- command
			}

		}
	}
}
