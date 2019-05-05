package connection

import (
	"net"
	"fmt"
	"strconv"
	"github.com/Redes-2019/userinterface"
)


type IrcClient struct {
	Socket   net.Conn
	UserInfo userinterface.User
	connInfo userinterface.ConnInfo
}

func NewClient(socket net.Conn, userInfo userinterface.User, connInfo userinterface.ConnInfo) *IrcClient {
	return &IrcClient{socket, userInfo, connInfo}
}


func OpenSocket(conn userinterface.ConnInfo) net.Conn {
	connTarget := conn.Servername + ":" + strconv.Itoa(conn.Port)
	fmt.Println("\nAbrindo o socket TCP para", connTarget)
	connSocket, err := net.Dial("tcp", connTarget)
	if err != nil {
		panic(err)
	} else {
		return connSocket
	}
}

// func
func (client *IrcClient) Connect(user userinterface.User ) {
	fmt.Println("Autenticando com o servidor")
	// Inicialmente manda PASS, se for necessário
	if client.connInfo.HasPasswd {
		passCmd := "PASS " + client.connInfo.Passwd + "\r\n"
		_, err := client.Socket.Write([]byte(passCmd))
		if err != nil {
			// fmt.Println("Erro enviando PASS")
			// fmt.Println("Enviado:", passCmd)
			// fmt.Println("Erro:", err)
			panic(err)
			return
		}
	}

	// Manda NICK
	client.UserInfo = user
	nickCmd := "NICK " + client.UserInfo.Nick + "\r\n"
	client.Socket.Write([]byte(nickCmd))
	_, err := client.Socket.Write([]byte(nickCmd))
	if err != nil {
		fmt.Println("Erro enviando NICK")
		fmt.Println("Enviado:", nickCmd)
		fmt.Println("Erro:", err)
		return
	}
	// Manda USER
	userCmd := ("USER " +
		client.UserInfo.Nick + " " +
		client.UserInfo.Hostname + " " +
		client.connInfo.Servername + " :" +
		client.UserInfo.Username + "\r\n")
	client.Socket.Write([]byte(userCmd))
	_, err = client.Socket.Write([]byte(userCmd))
	if err != nil {
		fmt.Println("Erro enviando USER")
		fmt.Println("Enviado:", userCmd)
		fmt.Println("Erro:", err)
		return
	}
}

func (client *IrcClient) Receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			client.Socket.Close()
			break
		}
		if length > 0 {
			fmt.Print("RECEIVED:" + string(message))
		}
	}
}

// Faz o handle da conexão e mandda a mensagem 
// pro servidor dependendo do comando

func HandleConnection(command []string) {
	fmt.Println("TODO: handler")
}