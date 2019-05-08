// Some color codes to print things nicely
package userinterface

// InfoTag is colored cyan for info messages
const InfoTag = cyan + "[info] " + reset

// OkTag is colored green for OK messages
const OkTag = green + "[ok] " + reset

// ErrorTag is colored black and red for errors
const ErrorTag = backRed + black + "[Fatal Error] " + reset

// WarnTag is colored yellow for warnings
const WarnTag = yellow + "[Warn] " + reset

// MsgOrigin is source of message
const MsgOrigin = blue

// CmdOrigin is command that gennerated the message
const CmdOrigin = magenta

// Resets print style
const Reset = "\u001b[0m"

// Foreground
const black = "\u001b[30m"
const red = "\u001b[31m"
const green = "\u001b[32m"
const yellow = "\u001b[33m"
const blue = "\u001b[34m"
const magenta = "\u001b[35m"
const cyan = "\u001b[36m"
const white = "\u001b[37m"
const reset = "\u001b[0m"

// Background
const backBlack = "\u001b[40m"
const backRed = "\u001b[41m"
const backGreen = "\u001b[42m"
const backYellow = "\u001b[43m"
const backBlue = "\u001b[44m"
const backMagenta = "\u001b[45m"
const backCyan = "\u001b[46m"
const backWhite = "\u001b[47m"

// Decorations
const bold = "\u001b[1m"
const underline = "\u001b[4m"
const reversed = "\u001b[7m"
