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
	fmt.Println(userinterface.InfoTag + "Iniciando Cliente.")
	// Pega informacões de usuário (nome, nick, server, etc)
	fmt.Println(userinterface.InfoTag + "Identificando usuário")
	user, conn := userinterface.ReadUserData()
	// Criando socket com o servidor
	connSocket := connection.OpenSocket(conn)
	// Criando cliente IRC
	client := connection.NewClient(connSocket, user, conn)
	go client.ListenServer()
	// Conectando cliente ao servidor
	for !<-client.ConnectSuccess {
		// Tenta autenticar com o servidor
		client.Connect()
		if <-client.NickInvalid {
			// Caso recebe um erro com o nick pede outro
			fmt.Print(userinterface.WarnTag + "Nick inválido. Escolha outro: ")
			nick, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			nick = strings.TrimRight(nick, "\n")
			client.UserInfo.Nick = nick
		} else if client.DeadSocket {
			// Neste ponto a conexão já foi fechada pelo servidor
			// Porque algum erro de autenticação aconteceu. e.g. senha incorreta
			return
		}
	}

	fmt.Println(userinterface.InfoTag + "Use /help to display available commands.")
	go listenUser(client)
	ok := true
	for ok {
		select {
		case msg, open := <-client.DataFromServer:
			// Mensagem foi recebida pelo servidor
			if !open {
				ok = false
			} else {
				fmt.Println(msg.PrintInfo, "<"+msg.Prefix+">", msg.Params)
			}
		case msg, _ := <-client.DataFromUser:
			// Usuário digitou alguma mensagem
			client.HandleConnection(msg)
		}
	}

	fmt.Println(userinterface.InfoTag + "Client disconnected. Terminating.")
}

// Pega input do usuário
func listenUser(client *connection.IrcClient) {
	channel := "> "
	// Entra em loop para ler os comandos do usuário
	for {
		command := userinterface.ReadCommand(channel)
		if command == nil {
			fmt.Println("[Input error] comando invalido")
		} else {
			validStructure, err := userinterface.VerifyStructure(command)
			if validStructure == false {
				fmt.Println(err)
			} else {
				client.DataFromUser <- command
			}
		}
	}
}
