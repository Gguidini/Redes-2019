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
	var validCommands [8]string
	validCommands[0] = "/join"
	validCommands[1] = "/list"
	validCommands[2] = "/quit"
	validCommands[3] = "/msg"
	validCommands[4] = "/part"
	validCommands[5] = "/help"
	validCommands[6] = "/topic"
	validCommands[7] = "/invite"
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
	if isValid {
		return parsedCommand
	}
	return nil
}

// VerifyStructure verifica a estrutura de comandos válidos.
// Evita que comandos com erro sintático sejam enviados ao servidor.
func VerifyStructure(command []string) (bool, string) {
	var result = false
	var err = ""
	switch strings.ToLower(command[0]) {

	case "/help":
		// Command: /help
		// Parameters: none
		// Displays available commands
		displayHelp()

	case "/join":
		// Command: /join
		// Parameters: <channel>{,<channel>} [<key>{,<key>}]
		if len(command) > 1 && len(command) <= 3 {
			channel := string(command[1])
			if (channel[0] == '#' || channel[0] == '&') && len(channel) > 1 {
				result = true
			} else {
				err = "Canal informado possui nome inválido. Nomes de canais iniciam com '#' ou '&'."
			}
		} else {
			err = `Número errado de parâmetros. Deveria ser 1 ou 2.
/join <channel>{,<channel>} [<key>{,<key>}]
Separe canais e keys APENAS por vírgula (sem espaço)`
		}

	case "/list":
		// Comand: /join
		// Parameters: [<channel>{,<channel>}]
		if len(command) <= 2 {
			result = true
		} else {
			err = `Número errado de parâmetros. Deveria ser 0 ou 1.
/list [<channel>{,<channel>}]
Separe canais APENAS por vírgula (sem espaço)`
		}

	case "/quit":
		if len(command) > 1 {
			message := string(command[1])
			if len(message) > 0 {
				result = true
			}
		} else {
			err = `Número errado de parâmetros. Deveria ser 1.
/quit <message>.
Mensagem de quit é obrigatória.`
		}

	case "/msg":
		// Parameters: <receiver>{,<receiver>} <text to be sent>
		if len(command) > 2 {
			target := string(command[1])
			message := string(command[2])
			if len(target) > 1 && len(message) > 1 {
				result = true
			}
		} else {
			err = `Número errado de parâmetros. Deveria ser 2.
/msg <receiver>{,<receiver>} <text to be sent>
Separe receivers APENAS por vírgula (sem espaço)`
		}

	case "/part":
		// Comand: /part
		// Parameters: <channel>{,<channel>}
		if len(command) == 2 {
			result = true
		} else {
			err = `Número errado de parâmetros. Deveria ser 1.
/part <channel>{,<channel>}
Separe canais APENAS por vírgula (sem espaço)`
		}

		// TODO: Command MODE
		// Comand: /mode
		// Parameters: <channel> {[+|-]|o|p|s|i|t|n|b|v} [<limit>] [<user>]	[<ban mask>]
		// Parameters: <nickname> {[+|-]|i|w|s|o}

	case "/topic":
		// Command: /topic
		// Parameters: <channel> [<topic>]
		if len(command) >= 2 {
			channel := command[1]
			if strings.Contains(channel, ",") {
				err = "Apenas 1 canal pode ser informado.\n/topic <channel> [<topic>]"
			} else {
				result = true
			}
		} else {
			err = `Número incorreto de parâmetros. Deveria ser 1 ou 2.
/topic <channel> [<topic>]`
		}

		// TODO: Command NAMES
		// Command: /names
		// Parameters: [<channel>{,<channel>}]

	case "/invite":
		// Command: /invite
		// Parameters: <nickname> <channel>
		if len(command) == 3 {
			result = true
		} else {
			err = `Número incorreto de parâmetros.
/invite <nickname> <channel>`
		}

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

	return result, err
}

func displayHelp() {
	availableCommands := []string{
		"/help - Displays available commands",
		"/join <channel>{,<channel>} [<key>{,<key>}] - Joins <channel using <key>.",
		"/part <channel>{,<channel>} - Leaves <channel>",
		"/list [<channel>{,<channel>}] - Displays visible channels, or info about <channel>.",
		"/quit <quit message> - Terminates connection with server. Message is mandatory.",
		"/msg <receiver>{,<receiver>} <text to be sent> - Sends message to <receiver>.",
		"/topic <channel> [<topic>] - Mostra o tópico de <channel>. Se <topic estiver presente, altera tópico de <channel> para <topic>.",
		"/invite <nickname> <channel> - Convida <nickname> para o canal <channel>.",
	}

	for _, help := range availableCommands {
		fmt.Println(help)
	}
}
