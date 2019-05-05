package connection

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/Redes-2019/userinterface"
)

// IrcClient é o cliente IRC
// Ele possui um socket TCP ligado ao servidor IRC
// E as informações de conexão e usuário.
// Também possui channels de input e output
type IrcClient struct {
	Socket         net.Conn
	UserInfo       userinterface.User
	connInfo       userinterface.ConnInfo
	DataFromServer chan string
	DataFromUser   chan []string
}

// NewClient retorna um novo IrcClient
func NewClient(socket net.Conn, userInfo userinterface.User, connInfo userinterface.ConnInfo) *IrcClient {
	return &IrcClient{socket, userInfo, connInfo, make(chan string, 100), make(chan []string, 100)}
}

// OpenSocket abre uma conexão com o servidor IRC
// Recebe as informações de conexão (userinterface.ConnInfo)
// Retorna o socket, caso tudo aconteça bem.
func OpenSocket(conn userinterface.ConnInfo) net.Conn {
	// Servidor:Porta
	connTarget := conn.Servername + ":" + strconv.Itoa(conn.Port)
	fmt.Println("\nAbrindo o socket TCP para", connTarget)
	connSocket, err := net.Dial("tcp", connTarget)
	if err != nil {
		panic(err)
	} else {
		// Quando uma conexão é aberta, o servidor tenta encontrar o usuário.
		// Vamos ler as mensagens de retorno para ver se o socket está ok
		// E para limpar o buffer de entrada
		// Como não passamos nenhum usuário, duas mensagens são enviadas:
		// :server_name NOTICE * :*** Looking up your username
		// :server_name NOTICE * :*** Could not find your username
		reader := bufio.NewReader(connSocket)
		msgLookUp, errLookUp := reader.ReadString('\n')
		_, errUserNotFound := reader.ReadString('\n')
		if errLookUp != nil {
			panic(errLookUp)
		} else if errUserNotFound != nil {
			panic(errUserNotFound)
		}
		server := strings.Fields(msgLookUp)[0]
		fmt.Println("Connection opened to", server[1:])
		return connSocket
	}
}

// Connect faz a autenticação com o servidor IRC com o qual temos um socket
// O cliente já tem o socket aberto com o servidor
// O cliente também já possui as informações de usuário e de conexão
// Autenticação é feita em 3 comandos:
// 1. PASS 2. NICK 3. USER
func (client *IrcClient) Connect() bool {
	fmt.Println("Autenticando com o servidor")
	// Inicialmente manda PASS, se for necessário
	if client.connInfo.HasPasswd {
		pass := passCmd(client.connInfo.Passwd)
		_, err := client.Socket.Write([]byte(pass))
		if err != nil {
			panic(err)
		}
	}

	// Manda NICK
	nick := nickCmd(client.UserInfo.Nick)
	_, err := client.Socket.Write([]byte(nick))
	if err != nil {
		panic(err)
	}
	// Manda USER
	user := userCmd(
		client.UserInfo.Nick,
		client.UserInfo.Hostname,
		client.connInfo.Servername,
		client.UserInfo.Username)
	_, err = client.Socket.Write([]byte(user))
	if err != nil {
		panic(err)
	}

	// Verifica se o servidor enviou as mensagens de Welcome.
	// Caso elas não tenham sido recebidas, autenticação falhou.
	reader := bufio.NewReader(client.Socket)
	msg, _ := reader.ReadString('\n')
	msgPieces := strings.Fields(strings.TrimRight(msg, CRLF))
	if msgPieces[1] != "001" {
		fmt.Println("Erro de Autenticação:", strings.Join(msgPieces[1:], " "))
		client.Socket.Close()
		return false
	}

	// Autenticação foi bem sucedida. Mostra mensagens de boas-vindas
	fmt.Println("Autenticação bem sucedida!")
	fmt.Println(strings.Join(msgPieces[3:], " "))
	for i := 0; i < 4; i++ {
		msg, _ = reader.ReadString('\n')
		msgPieces = strings.Fields(strings.TrimRight(msg, CRLF))
		fmt.Println(strings.Join(msgPieces[3:], " "))
	}
	return true
}

// ListenServer recebe as mensagens do servidor e as passa para o display.
// Caso receba uma mensagem de ERROR ou KILL fecha o socket
// Caso receba uma mensagem PING, envia o PONG
func (client *IrcClient) ListenServer() {
	readSocket := bufio.NewReader(client.Socket)
	for {
		message, err := readSocket.ReadString('\n')
		message = strings.TrimRight(message, "\r\n")
		if err != nil {
			fmt.Println("Erro lendo so socket:", err)
			fmt.Println("Fechando conexão")
			client.Socket.Close()
			close(client.DataFromServer)
			break
		}
		// Mensages que iniciam com prefixo não são erros
		// Então são mostradas
		if message[0] == ':' {
			client.DataFromServer <- message[1:]
		}
		// Mensagens que não iniciam com prefixo podem ser erros ou pings
		fields := strings.Fields(message)
		if fields[0] == "PING" {
			// Mensagens de PING devem ser respondidas para manter o canal aberto.
			// A resposta é um PONG
			reply := pongCmd(fields[1])
			client.Socket.Write([]byte(reply))
		} else if fields[0] == "ERROR" || fields[0] == "KILL" {
			// Mensagens de ERROR significam que algo deu errado e o servidor fechou a conexão
			// KILL significa que a conexão foi fechada por algum operador
			// Logo, o cliente precisará se reconectar.
			fmt.Println("Fatal Error:", message)
			client.Socket.Close()
			close(client.DataFromServer)
			break
		}
	}
	fmt.Println("Stopped listening.")
}

// HandleConnection envia comandos para o servidor pelo socket.
// Os comandos são inseridos pelo usuário.
func (client *IrcClient) HandleConnection(command []string) {
	var cmdToSend string
	switch command[0] {
	case "/join":
		if len(command) == 3 {
			// Channel e Key
			cmdToSend = joinCmd(command[1], command[2])
		} else {
			// Just Channel
			cmdToSend = joinCmd(command[1], "")
		}
	case "/part":
		cmdToSend = partCmd(command[1])
	case "/quit":
		cmdToSend = strings.Join(command[1:], " ")
		cmdToSend = quitCmd(cmdToSend)
	case "/list":
		if len(command) >= 2 {
			// Channel e Key
			cmdToSend = listCmd(command[1])
		} else {
			// Just Channel
			cmdToSend = listCmd("")
		}
	case "/msg":
		cmdToSend = msgCmd(command[1], command[2:])
	}
	fmt.Println("[Sending]", cmdToSend)
	client.Socket.Write([]byte(cmdToSend))
}
