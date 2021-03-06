// Package connection gerencia a conexão com o servidor
package connection

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/Redes-2019/userinterface"
)

// IrcClient é o cliente IRC
// Ele possui um socket TCP ligado ao servidor IRC
// E as informações de conexão e usuário.
// Também possui channels de input e output
// E flags de estado
type IrcClient struct {
	Socket         net.Conn
	UserInfo       userinterface.User
	connInfo       userinterface.ConnInfo
	DataFromServer chan Message
	DataFromUser   chan []string
	isAway         bool
	DeadSocket     bool
	ConnectSuccess chan bool
	NickInvalid    chan bool
}

// NewClient retorna um novo IrcClient
func NewClient(socket net.Conn, userInfo userinterface.User, connInfo userinterface.ConnInfo) *IrcClient {
	c := &IrcClient{socket, userInfo, connInfo, make(chan Message, 100), make(chan []string, 100), false, false, make(chan bool, 1), make(chan bool, 1)}
	c.ConnectSuccess <- false // Inicia conexão como não feita
	return c
}

// OpenSocket abre uma conexão com o servidor IRC
// Recebe as informações de conexão (userinterface.ConnInfo)
// Retorna o socket, caso tudo aconteça bem.
func OpenSocket(conn userinterface.ConnInfo) net.Conn {
	// Servidor:Porta
	connTarget := conn.Servername + ":" + strconv.Itoa(conn.Port)
	fmt.Println(userinterface.InfoTag+"Abrindo o socket TCP para", connTarget)
	connSocket, err := net.Dial("tcp", connTarget)
	if err != nil {
		panic(err)
	}

	fmt.Println(userinterface.OkTag + "Conexão bem sucedida!")
	return connSocket
}

// Connect faz a autenticação com o servidor IRC com o qual temos um socket
// O cliente já tem o socket aberto com o servidor
// O cliente também já possui as informações de usuário e de conexão
// Autenticação é feita em 3 comandos:
// 1. PASS 2. NICK 3. USER
func (client *IrcClient) Connect() {
	fmt.Println(userinterface.InfoTag + "Autenticando com o servidor")
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
}

// ListenServer recebe as mensagens do servidor e as passa para o display.
// Caso receba uma mensagem de ERROR ou KILL fecha o socket
// Caso receba uma mensagem PING, envia o PONG
// Verifica alguns erros de autenticação
func (client *IrcClient) ListenServer() {
	// Socket, buffer de leitura
	readSocket := bufio.NewReader(client.Socket)
	for {
		// Lê mensagem recebida pelo Socket
		message, err := readSocket.ReadString('\n')
		if err != nil {
			fmt.Println(userinterface.ErrorTag, err)
			fmt.Println(userinterface.WarnTag + "Fechando conexão")
			// Clean up
			client.DeadSocket = true
			client.ConnectSuccess <- false
			client.NickInvalid <- false
			client.Socket.Close()
			close(client.DataFromServer)
			break
		}
		// Formata a mensagem recebida
		parsedMsg := parseMessage(message)
		// Mensages que iniciam com prefixo não são erros, então são mostradas.
		// A mensagem de PING é para controle, não precisa ser mostrada
		if parsedMsg.Cmd != "PING" && parsedMsg.Cmd != "ERROR" {
			client.DataFromServer <- parsedMsg
		}

		// Verifies if message has a numeric code
		if parsedMsg.Cmd[0] >= ' ' && parsedMsg.Cmd[0] <= '9' {
			// Check for specific responses to set flags
			switch parsedMsg.Cmd {
			// Some Nick Error - Auth failed
			case erroneusNick, errNickCollision, errNickUsed:
				client.ConnectSuccess <- false
				client.NickInvalid <- true
			// Welcome Messages - Auth was completed!
			case welcomeHeader1:
				client.ConnectSuccess <- true
				client.NickInvalid <- false
			}
		}
		// Mensagens que não iniciam com prefixo podem ser erros ou pings
		if parsedMsg.Cmd == "PING" {
			// Mensagens de PING devem ser respondidas para manter o canal aberto.
			// A resposta é um PONG
			reply := pongCmd(parsedMsg.Params)
			client.Socket.Write([]byte(reply))
		} else if parsedMsg.Cmd == "ERROR" || parsedMsg.Cmd == "KILL" {
			// Mensagens de ERROR significam que algo deu errado e o servidor fechou a conexão
			// KILL significa que a conexão foi fechada por algum operador
			// Logo, o cliente precisará se reconectar.
			fmt.Println(userinterface.ErrorTag, parsedMsg.Params)
			client.Socket.Close()
			close(client.DataFromServer)
			client.DeadSocket = true
			client.ConnectSuccess <- false
			client.NickInvalid <- false
			return
		}
	}
	fmt.Println(userinterface.WarnTag + "Stopped listening.")
}

// HandleConnection envia comandos para o servidor pelo socket.
// Os comandos são inseridos pelo usuário.
func (client *IrcClient) HandleConnection(command []string) {
	var cmdToSend string
	// Verifica qual comando será enviado
	switch strings.ToLower(command[0]) {
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
			cmdToSend = listCmd(command[1])
		} else {
			cmdToSend = listCmd("")
		}
	case "/msg":
		cmdToSend = msgCmd(command[1], command[2:])
	case "/topic":
		cmdToSend = topicCmd(command[1], strings.Join(command[2:], " "))
	case "/invite":
		cmdToSend = inviteCmd(command[1], command[2])
	case "/names":
		if len(command) == 2 {
			cmdToSend = namesCmd(command[1])
		} else {
			cmdToSend = namesCmd("")
		}
	case "/ison":
		cmdToSend = isonCmd(command[1:])
	case "/away":
		if len(command) == 1 && !client.isAway {
			// Tentado enviar desAWAY command sem estar AWAY
			return
		} else if len(command) == 1 && client.isAway {
			client.isAway = true
			cmdToSend = awayCmd(nil)
		} else {
			cmdToSend = awayCmd(command[1:])
		}
	case "/who":
		if len(command) == 3 {
			cmdToSend = whoCmd(command[1], command[2])
		} else if len(command) == 2 {
			cmdToSend = whoCmd(command[1], "")
		} else {
			cmdToSend = whoCmd("", "")
		}
	case "/whois":
		cmdToSend = whoisCmd(command[1])
	case "/mode":
		cmdToSend = modeCmd(command[1:])
	case "/kick":
		if len(command) == 3 {
			cmdToSend = kickCmd(command[1], command[2], nil)
		} else {
			cmdToSend = kickCmd(command[1], command[2], command[3:])
		}
	case "/clear":
		fmt.Println("test")
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	}

	client.Socket.Write([]byte(cmdToSend))
}
