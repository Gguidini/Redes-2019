package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/Redes-2019/connection"
	"github.com/Redes-2019/userinterface"
	"github.com/Redes-2019/tui"
)

func main() {
	fmt.Println("==> Iniciando Cliente.")
	// Pega informacões de usuário
	fmt.Println("==> Identificando usuário")
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
			fmt.Print("Nick inválido. Escolha outro: ")
			nick, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			nick = strings.TrimRight(nick, "\n")
			client.UserInfo.Nick = nick
		} else {
			// Neste ponto a conexão já foi fechada
			return
		}
	}

	fmt.Println("[info] Use /help to display available commands.")
	go listenUser(client)
	ok := true
	for ok {
		select {
		case msg, open := <-client.DataFromServer:
			if !open {
				ok = false
			} else {
				fmt.Println(msg.PrintInfo, "<"+msg.Prefix+">", msg.Params)
			}
		case msg, _ := <-client.DataFromUser:
			client.HandleConnection(msg)
		}
	}

	fmt.Println("[info] Client disconnected. Terminating.")
}

func listenUser(client *connection.IrcClient) {
	// Inicia as funcões da TUI
	tui.TuiHandler(client)
}
