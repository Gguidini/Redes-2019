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

func userCmd(nick, host, server, user string) string {
	cmd := "USER " +
		nick + " " +
		host + " " +
		server + " :" +
		user + CRLF
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

func listCmd(channel string) string {
	if channel == "" {
		return "LIST" + CRLF
	}
	return "LIST " + channel + CRLF
}

func topicCmd(channel, topic string) string {
	cmd := "TOPIC " + channel
	if topic != "" {
		cmd += " :" + topic
	}
	cmd += CRLF
	return cmd
}

func inviteCmd(nick, channel string) string {
	return "INIVTE " + nick + " " + channel + CRLF
}

func namesCmd(channel string) string {
	cmd := "NAMES"
	if channel != "" {
		cmd += " " + channel
	}
	cmd += CRLF
	return cmd
}
