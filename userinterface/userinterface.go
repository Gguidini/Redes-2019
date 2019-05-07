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
	fmt.Print("Porta da conexão (deixe branco para usar default): ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimRight(port, "\n")
	if port == "" {
		port = "6667"
	} else {
		port = strings.TrimRight(port, "\n")
	}
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
	var validCommands [16]string
	validCommands[0] = "/join"
	validCommands[1] = "/list"
	validCommands[2] = "/quit"
	validCommands[3] = "/msg"
	validCommands[4] = "/part"
	validCommands[5] = "/help"
	validCommands[6] = "/topic"
	validCommands[7] = "/invite"
	validCommands[8] = "/names"
	validCommands[9] = "/ison"
	validCommands[10] = "/away"
	validCommands[11] = "/who"
	validCommands[12] = "/mode"
	validCommands[13] = "/whois"
	validCommands[14] = "/kick"
	validCommands[15] = "/clear"
	for _, item := range validCommands {
		if item == command || len(command) == 0{
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
	fmt.Print("> ")
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

	case "/names":
		// Command: /names
		// Parameters: [<channel>{,<channel>}]
		if len(command) <= 2 {
			result = true
		} else {
			err = `Número incorreto de parâmetros. Deveria ser 0 ou 1.
/names [<channel>{,<channel>}]
Separe canais APENAS por vírgula (sem espaço)`
		}

	case "/invite":
		// Command: /invite
		// Parameters: <nickname> <channel>
		if len(command) == 3 {
			result = true
		} else {
			err = `Número incorreto de parâmetros. Deveria ser 2.
/invite <nickname> <channel>`
		}

	case "/who":
		// Command: /who
		// Parameters: [<name> [<o>]]
		if len(command) > 1 && len(command) <= 3 {
			if len(command) == 3 && command[2] == "o" {
				result = true
			} else if len(command) == 2 {
				result = true
			} else {
				err = "Último argumento só pode ser 'o'. /who [<mask> ['o']]."
			}
		} else {
			err = `Número de argumentos inválido. Deve ser 0 até 2.
/who [<mask> ['o']].`
		}

	case "/away":
		// Command: /away
		// Parameters: [message]
		result = true

	case "/ison":
		// Command: /ison
		// Parameters: <nickname>{<space><nickname>}
		if len(command) > 1 {
			result = true
		} else {
			err = `Número incorreto de parâmetros. Deveria ser pelo menos 1.
/ison <nickname>{<space><nickname>}`
		}

	case "/whois":
		// Command: /whois
		// Parameters: <nickmask>{,<nichmask>}
		if len(command) == 2 {
			result = true
		} else {
			err = `Número incorreto de parâmetro. Deve ser 2.
/whois <nickmask>{,<nichmask>}.
Separe <nickamsk> APENAS por vírgula (sem espaço)`
		}

	case "/kick":
		// Command: /kick
		// Parameters: <channel> <user> [<comment>]
		if len(command) >= 3 {
			if command[1][0] == '#' || command[1][0] == '&' {
				result = true
			} else {
				err = "Primeiro argumento deve ser um canal"
			}
		} else {
			err = `Número inválido de parâmetros. devem ser ao menos 2.
/kick <channel> <user> [<comment>].
Necessário privilégios para o comando funcionar.
`
		}
	case "/mode":
		// Comand: /mode
		// Parameters: <channel> {[+|-]|o|p|s|i|t|n|b|v|k} [<limit>] [<user>]	[<ban mask>] [<key>]
		// Parameters: <nickname> {[+|-]|i|w|s|o}
		if len(command) >= 2 && len(command) < 5 {
			if len(command) == 2 {
				result = true
			} else if command[1][0] == '#' || command[1][0] == '&' {
				// Mode de canal
				if command[2][0] == '+' || command[2][0] == '-' {
					if strings.ContainsAny(command[2], "obvkl") && len(command) >= 4 {
						result = true
					} else if !strings.ContainsAny(command[2], "obvkl") && len(command) >= 4 {
						err = "Muitos argumentos para este comando."
					} else {
						result = true
					}
				} else {
					err = "Indique se vai ativar (+) ou desativar (-) as opções."
				}
			} else {
				if len(command) == 3 {
					if command[2][0] == '+' || command[2][0] == '-' {
						result = true
					} else {
						err = "Indique se vai ativar (+) ou desativar (-) as opções."
					}
				} else {
					err = "Número inválido de argumentos. Deveriam ser 3."
				}
			}
		} else {
			err = `Número errado de parâmetros. Deve ser entre 1 e 5.
/mode <channel> {[+|-]|o|p|s|i|t|n|b|v} [<limit>] [<user>]	[<ban mask>]
/mode <nickname> {[+|-]|i|w|s|o}
[limit] usado com flag l, [user] com flags o,v, [key] com flag k, [banmask] com flag b.`
		}
	case "/clear":
		result = true
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
		"/names [<channel>{,<channel>}] - Lista usuários disponíveis no canal. Se nenhum canal é informado, lista todos os usuários visíveis.",
		"/ison <nickname>{<space><nickname>} - Verifica se <nickname> está ON",
		"/away [message] - Define usuário como AWAY, se tiver uma mensagem, ou cancela AWAY, se não houver mensagem.",
		"/who [<mask> ['o']] - Busca informações sobre qualquer usuário que seja match com a mask (use regex). Se a opção 'o' estiver ativada, retorna somente sobre operators.",
		"/whois <nickmask>{,<nichmask>} - Mostra mais informações sobre determinado usuário",
		"/kick <channel> <user> [<comment>] - Exclui <user> de <chanel>. Necessita privilégios.",
		"/clear - Limpa a tela",
		"",
		`/mode <channel> {[+|-]|o|p|s|i|t|n|b|v|k} [<limit>] [<user>] [<ban mask>] [<key>] - Altera o mode de um canal, ou lista os modes dele se não houver flags. Algumas opções precisam de privilégios para serem aceitas. (+) Ativa flag, (-) desativa flag. Flags são:
		o - give/take channel operator privileges;
    	p - private channel flag;
    	s - secret channel flag;
    	i - invite-only channel flag;
        t - topic settable by channel operator only flag;
    	n - no messages to channel from clients on the outside;
    	m - moderated channel;
		l - set the user limit to channel;
		b - set a ban mask to keep users out;
    	v - give/take the ability to speak on a moderated channel;
		k - set a channel key (password).`,
		"",
		`/mode <nickname> {[+|-]|i|w|s|o} - Altera modo de usuário, ou lista os modes dele se não houver flag. (+) Ativa flag, (-) desativa flag. Algumas opções precisam de privilégios. +o é sempre ignorado. Flags são:
		i - marks a users as invisible;
    	s - marks a user for receipt of server notices;
    	w - user receives wallops;
    	o - operator flag.`,
	}

	for _, help := range availableCommands {
		fmt.Println(help)
	}
}
