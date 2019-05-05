package main
import (
	"fmt"
)

import "github.com/Redes-2019/userinterface"

func main() {
	user := new(userinterface.User)
	user = userinterface.ReadUserData()
	fmt.Print((*user).Username)
}

// func