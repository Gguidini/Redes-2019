package connection

import "strings"

const CRLF = "\r\n"

func pongCmd(target string) string {
	return "PONG " + target + CRLF
}

func passCmd(passwd string) string {
	return "PASS " + passwd + CRLF
}

func nickCmd(nick string) string {
	return "NICK " + nick + CRLF
}

func userCmd(nick, server, host string, user []string) string {
	cmd := "USER " +
		nick + " " +
		host + " " +
		server + " :" +
		strings.Join(user, " ") +
		CRLF
	return cmd
}

func quitCmd(quitMsg string) string {
	return "QUIT :" + quitMsg + CRLF
}

func joinCmd(channel string, key string) string {
	cmd := "JOIN " + channel
	if key != "" {
		cmd += " " + key
	}
	cmd += CRLF
	return cmd
}

func partCmd(channel string) string {
	return "PART " + channel + CRLF
}

func msgCmd(receiver string, message []string) string {
	return "PRIVMSG " + receiver + " :" + strings.Join(message, " ") + CRLF
}
