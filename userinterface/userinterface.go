package userinterface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// User define as informações de identificação de um cliente IRC.
// As informações são:
// Username: nome real do usuário
// Nick: nome utilizado como identificador
// Hostname: máquina em que o cliente está rodando
type User struct {
	Username string
	Nick     string
	Hostname string
}

// ConnInfo define as informações de conexão de um cliente IRC.
// As informações são:
// Servername: servidor IRC no qual o cliente vai se conectar
// HasPasswd: se aquela conexão precisa de uma senha ou não
// Passwd: senha da conexão se HasPasswd == true, "" se HasPasswd == false
type ConnInfo struct {
	Servername string
	Port       int
	Passwd     string
	HasPasswd  bool
}

// ReadUserData lê as informações do usuário utilizando o cliente IRC
// As informações são aquelas da struct User e ConnInfo
func ReadUserData() (User, ConnInfo) {
	reader := bufio.NewReader(os.Stdin)

	// Username
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimRight(username, "\n")

	// Nickname
	fmt.Print("Nick: ")
	nick, _ := reader.ReadString('\n')
	nick = strings.TrimRight(nick, "\n")

	// Server to connect
	fmt.Print("Remote server name: ")
	server, _ := reader.ReadString('\n')
	server = strings.TrimRight(server, "\n")

	// Connection port
	fmt.Print("Porta da conexão: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimRight(port, "\n")
	numPort, _ := strconv.Atoi(port)

	// Connection Password, if any
	fmt.Print("O servidor possui uma senha? [s/n] ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimRight(pass, "\n")
	pass = strings.TrimLeft(pass, "\n")
	passFlag := false

	if strings.ToLower(pass) == "s" {
		passFlag = true
		fmt.Print("Senha: ")
		pass, _ = reader.ReadString('\n')
		pass = strings.TrimRight(pass, "\n")
	} else {
		passFlag = false
		pass = ""
	}

	host, err := os.Hostname()
	if err != nil {
		host = "Unkown"
	}

	newUser := User{username, nick, host}
	newConnInfo := ConnInfo{server, numPort, pass, passFlag}
	return newUser, newConnInfo
}

// Faz o parse do comando em um array de strings

func parseCommand(command string) []string {
	parsedString := strings.Split(command, " ")
	return parsedString
}

// Valida o comando recebido
func validateCommand(command string) bool {
	var validCommands [5]string
	validCommands[0] = "/join"
	validCommands[1] = "/list"
	validCommands[2] = "/quit"
	validCommands[3] = "/msg"
	validCommands[4] = "/part"
	for _, item := range validCommands {
		if item == command {
			return true
		}
	}
	return false
}

// Lê o comando recebido da main
func ReadCommand(channel string) []string {
	reader := bufio.NewReader(os.Stdin)

	// Imprime o canal que o usuário se encontra,
	// ou nenhum se ele não está num canal
	fmt.Print(channel)
	command, _ := reader.ReadString('\n')
	command = strings.TrimRight(command, "\r\n")
	parsedCommand := parseCommand(command)
	isValid := validateCommand(strings.ToLower(parsedCommand[0]))
	if isValid == true {
		return parsedCommand
	}
	return nil
}

// Verifica a estrutura de cada comando
// /join <#channel>
// /list
// /quit <message>
// /msg <#channel>|<user> <message>
func VerifyStructure(command []string) bool {
	var result = false
	switch command[0] {

	// TODO: Add help command to list available commands
	// Command: /help
	// Parameters: none

	case "/join":
		// TODO: Fix join parsing to accept multiple channels and keys. Must be comma separated, no space
		// Command: /join
		// Parameters: <channel>{,<channel>} [<key>{,<key>}]
		if len(command) > 1 {
			channel := string(command[1])
			if (channel[0] == '#' || channel[0] == '&') && len(channel) > 1 {
				result = true
			}
		}

	case "/list":
		// TODO: Accept multiple channels. Must be comma separated, no space
		// Parameters: [<channel>{,<channel>}]
		result = true

	case "/quit":
		if len(command) > 1 {
			message := string(command[1])
			if len(message) > 0 {
				result = true
			}
		}

	case "/msg":
		// TODO: Accept multiple receivers. Must be comma separated, no space
		// Parameters: <receiver>{,<receiver>} <text to be sent>
		if len(command) > 2 {
			target := string(command[1])
			message := string(command[2])
			if len(target) > 1 && len(message) > 1 {
				result = true
			}
		}
	case "/part":
		// TODO: Part parsing has to accept multiple channels. Must be comma separated, no space
		// Comand: /part
		// Parameters: <channel>{,<channel>}
		if len(command) == 2 {
			result = true
		}

		// TODO: Command MODE
		// Comand: /mode
		// Parameters: <channel> {[+|-]|o|p|s|i|t|n|b|v} [<limit>] [<user>]	[<ban mask>]
		// Parameters: <nickname> {[+|-]|i|w|s|o}

		// TODO: Command TOPIC
		// Command: /topic
		// Parameters: <channel> [<topic>]

		// TODO: Command NAMES
		// Command: /names
		// Parameters: [<channel>{,<channel>}]

		// TODO: Command INVITE
		// Command: /invite
		// Parameters: <nickname> <channel>

		// TODO: Command WHO
		// Command: /who
		// Parameters: [<name> [<o>]]

		// TODO: Command AWAY
		// Command: /away
		// Parameters: [message]

		// TODO: Command ISON
		// Command: /ison
		// Parameters: <nickname>{<space><nickname>}
	}

	return result
}
