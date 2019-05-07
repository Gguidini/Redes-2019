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

func isonCmd(nicks []string) string {
	return "ISON " + strings.Join(nicks, " ") + CRLF
}

func awayCmd(msg []string) string {
	if msg == nil {
		return "AWAY" + CRLF
	}
	return "AWAY :" + strings.Join(msg, " ") + CRLF
}

func whoCmd(mask, o string) string {
	cmd := "WHO"
	if mask != "" {
		cmd += " " + mask
	}
	if o != "" {
		cmd += " " + o
	}
	cmd += CRLF
	return cmd
}

func whoisCmd(nicks string) string {
	return "WHOIS " + nicks + CRLF
}
func modeCmd(options []string) string {
	return "MODE " + strings.Join(options, " ") + CRLF
}

func kickCmd(channel, user string, comment []string) string {
	if comment == nil {
		return "KICK " + channel + " " + user + CRLF
	}
	return "KICK " + channel + " " + user + " " + strings.Join(comment, " ") + CRLF
}
