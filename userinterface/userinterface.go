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
