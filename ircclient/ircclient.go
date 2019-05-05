package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/Redes-2019/userinterface"
)

type ircClient struct {
	socket   net.Conn
	userInfo userinterface.User
	connInfo userinterface.ConnInfo
}

func main() {
	fmt.Println("Iniciando Cliente.")
	fmt.Println("Identificando usuário")
	user, conn := userinterface.ReadUserData()

	connTarget := conn.Servername + ":" + strconv.Itoa(conn.Port)
	fmt.Println("\nAbrindo o socket TCP para", connTarget)
	connSocket, err := net.Dial("tcp", connTarget)

	if err != nil {
		fmt.Println("Error connecting:", err)
	}
	client := &ircClient{socket: connSocket, userInfo: user, connInfo: conn}

	go client.receive()
	client.connect()

	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		strings.TrimRight(message, "\n")
		message += "\r\n"
		fmt.Println("Sending to server:", message)
		connSocket.Write([]byte(message))
	}

}

// func
func (client *ircClient) connect() {
	fmt.Println("Autenticando com o servidor")
	// Inicialmente manda PASS, se for necessário
	if client.connInfo.HasPasswd {
		passCmd := "PASS " + client.connInfo.Passwd + "\r\n"
		_, err := client.socket.Write([]byte(passCmd))
		if err != nil {
			fmt.Println("Erro enviando PASS")
			fmt.Println("Enviado:", passCmd)
			fmt.Println("Erro:", err)
			return
		}
	}

	// Manda NICK
	nickCmd := "NICK " + client.userInfo.Nick + "\r\n"
	client.socket.Write([]byte(nickCmd))
	_, err := client.socket.Write([]byte(nickCmd))
	if err != nil {
		fmt.Println("Erro enviando NICK")
		fmt.Println("Enviado:", nickCmd)
		fmt.Println("Erro:", err)
		return
	}
	// Manda USER
	userCmd := ("USER " +
		client.userInfo.Nick + " " +
		client.userInfo.Hostname + " " +
		client.connInfo.Servername + " :" +
		client.userInfo.Username + "\r\n")
	client.socket.Write([]byte(userCmd))
	_, err = client.socket.Write([]byte(userCmd))
	if err != nil {
		fmt.Println("Erro enviando USER")
		fmt.Println("Enviado:", userCmd)
		fmt.Println("Erro:", err)
		return
	}
}

func (client *ircClient) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
		}
	}
}
