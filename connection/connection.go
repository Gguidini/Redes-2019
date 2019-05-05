package connection

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/Redes-2019/userinterface"
)

type IrcClient struct {
	Socket   net.Conn
	UserInfo userinterface.User
	connInfo userinterface.ConnInfo
	data     chan string
}

func NewClient(socket net.Conn, userInfo userinterface.User, connInfo userinterface.ConnInfo) *IrcClient {
	return &IrcClient{socket, userInfo, connInfo, make(chan string, 100)}
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
func (client *IrcClient) Connect(user userinterface.User) {
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
	readSocket := bufio.NewReader(client.Socket)
	for {
		message, err := readSocket.ReadString('\n')
		message = strings.TrimRight(message, "\r\n")
		if err != nil {
			fmt.Println("Erro lendo so socket:", err)
			fmt.Println("Fechando conexão")
			client.Socket.Close()
			close(client.data)
			break
		}
		// Mensages que iniciam com prefixo não são erros
		// Então são mostradas
		if message[0] == ':' {
			fmt.Println("RECEIVED:", message)
		}
		// Mensagens que não iniciam com prefixo podem ser erros ou pings
		fields := strings.Fields(message)
		if fields[0] == "PING" {
			// Mensagens de PING devem ser respondidas para manter o canal aberto.
			// A resposta é um PONG
			fmt.Println("Received PING. Responding.")
			reply := "PONG " + fields[1] + "\r\n"
			client.Socket.Write([]byte(reply))
		} else if fields[0] == "ERROR" {
			// Mensagens de ERROR significam que algo deu errado e o servidor fechou a conexão
			// Logo, o cliente precisará se reconectar.
			fmt.Println("Fatal Error:", message)
			client.Socket.Close()
			close(client.data)
			break
		}
	}
	fmt.Println("Stopped listening.")
}

// Faz o handle da conexão e mandda a mensagem
// pro servidor dependendo do comando
func HandleConnection(command []string) {
	fmt.Println("TODO: handler")
}
