package connection

import "strings"

// Message define uma mensagem que vem do servidor.
// Definição das mensagens segundo RFC 1459
// https://tools.ietf.org/html/rfc1459#section-2.3.1
// <message>  ::= [':' <prefix> <SPACE> ] <command> <params> <crlf>
// <prefix>   ::= <servername> | <nick> [ '!' <user> ] [ '@' <host> ]
// <command>  ::= <letter> { <letter> } | <number> <number> <number>
// <SPACE>    ::= ' ' { ' ' }
// <params>   ::= <SPACE> [ ':' <trailing> | <middle> <params> ]
// <middle>   ::= <Any *non-empty* sequence of octets not including SPACE
//                or NUL or CR or LF, the first of which may not be ':'>
// <trailing> ::= <Any, possibly *empty*, sequence of octets not including
//                  NUL or CR or LF>
// <crlf>     ::= CR LF
type Message struct {
	Prefix    string
	Cmd       string
	Params    string
	PrintInfo string
}

// parseMessage quebra uma mensagem recebida do servidor em campos.
// Campo PrintInfo é para informar ao usuário depois.
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
	case mOTDbody, mOTDhead, mOTDtail, mOTDmissing:
		parsedMsg.PrintInfo = "[MOTD]"
	case whoEnd, whoRpl:
		parsedMsg.PrintInfo = "[Who]"
	case awayOn, awayOff:
		parsedMsg.PrintInfo = "[Away]"
	case whoisChan, whoisEnd, whoisIdle, whoisOper, whoisServer, whoisUser:
		parsedMsg.PrintInfo = "[Whois]"
	case ison:
		parsedMsg.PrintInfo = "[IsOn]"
	case names, namesEnd:
		parsedMsg.PrintInfo = "[Names]"
	case topic, topicNo:
		parsedMsg.PrintInfo = "[Topic]"
	case userMode, chanMode:
		parsedMsg.PrintInfo = "[Mode]"
	case listBody, listEnd, listHead:
		parsedMsg.PrintInfo = "[List]"
	case inviteOK:
		parsedMsg.PrintInfo = "[Invite]"
		parsedMsg.Params += " Invite successful!"
	case "NOTICE":
		parsedMsg.PrintInfo = "[Notice]"
	case "ERROR":
		parsedMsg.PrintInfo = "[ERROR]"
	case "KILL":
		parsedMsg.PrintInfo = "[KILL]"
	}

	return parsedMsg
}
