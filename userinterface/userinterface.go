package userinterface

import (
	"fmt"
	"bufio"
	"os"
)

// Atributos com letra maíuscula podem ser acessados de fora
type User struct {
	Username string
	Nick string
	Password string
}

func ReadUserData() *User {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("nick: ")
	nick, _ := reader.ReadString('\n')
	fmt.Print("senha: ")
	password, _ := reader.ReadString('\n')

	newUser := User{username, nick, password}
	return &newUser
}

