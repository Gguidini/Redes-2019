package connection

import "strings"

// Alguns códigos de erro que podem acontecer quando registrando
const erroneusNick = "432"
const errNickUsed = "433"
const errNickCollision = "436"
const welcomeHeader1 = "001"
const welcomeHeader2 = "002"
const welcomeHeader3 = "003"
const welcomeHeader4 = "004"
const welcomeHeader5 = "005"

type Message struct {
	Prefix    string
	Cmd       string
	Params    string
	PrintInfo string
}

func parseMessage(rawMsg string) Message {
	rawMsg = strings.TrimRight(rawMsg, "\r\n")
	msgPieces := strings.Fields(rawMsg)
	var parsedMsg Message
	if msgPieces[0][0] == ':' {
		// Message has a prefix
		// <prefix>   ::= <servername> | <nick> [ '!' <user> ] [ '@' <host> ]
		// Extracting <user>@<host> option, if present
		// Or <server> or <nick>
		var userAndHost string
		if i := strings.Index(msgPieces[0], "!"); i > -1 {
			userAndHost = msgPieces[0][i+1:]
			if userAndHost[0] == '~' {
				userAndHost = userAndHost[1:]
			}
		} else {
			userAndHost = msgPieces[0][1:]
		}
		parsedMsg.Prefix = userAndHost
		parsedMsg.Cmd = msgPieces[1]
		if len(msgPieces) > 2 {
			parsedMsg.Params = strings.Join(msgPieces[2:], " ")
			parsedMsg.Params = strings.TrimLeft(parsedMsg.Params, ":")
		}
	} else {
		parsedMsg.Prefix = ""
		parsedMsg.Cmd = msgPieces[0]
		if len(msgPieces) > 1 {
			parsedMsg.Params = strings.Join(msgPieces[1:], " ")
			parsedMsg.Params = strings.TrimLeft(parsedMsg.Params, ":")
		}
	}

	switch parsedMsg.Cmd {
	// Some Nick Error
	case erroneusNick, errNickCollision, errNickUsed:
		parsedMsg.PrintInfo = "[Warn]"
	// Welcome Messages
	case welcomeHeader1, welcomeHeader2, welcomeHeader3, welcomeHeader4, welcomeHeader5:
		parsedMsg.PrintInfo = "[Welcome]"
	case "NOTICE":
		parsedMsg.PrintInfo = "[Notice]"
	case "ERROR":
		parsedMsg.PrintInfo = "[ERROR]"
	}

	return parsedMsg
}
