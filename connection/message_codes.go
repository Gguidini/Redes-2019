// Definições de códigos de mensagens
package connection

// Códigos numéricos de algumas mensagens.
// Servem para identificar o comando que as originou
// NICK
const erroneusNick = "432"
const errNickUsed = "433"
const errNickCollision = "436"

// WELCOME
const welcomeHeader1 = "001"
const welcomeHeader2 = "002"
const welcomeHeader3 = "003"
const welcomeHeader4 = "004"
const welcomeHeader5 = "005"

// MOTD
const mOTDhead = "375"
const mOTDbody = "372"
const mOTDtail = "376"
const mOTDmissing = "422"

// WHO
const whoRpl = "352"
const whoEnd = "315"

// AWAY
const awayOff = "305"
const awayOn = "306"

// WHOIS
const whoisUser = "311"
const whoisServer = "312"
const whoisOper = "313"
const whoisIdle = "317"
const whoisEnd = "318"
const whoisChan = "319"

// ISON
const ison = "303"

// NAMES
const names = "353"
const namesEnd = "366"

// TOPIC
const topic = "332"
const topicNo = "331"

// MODE
const chanMode = "324"
const userMode = "221"

// LIST
const listHead = "321"
const listBody = "322"
const listEnd = "323"

// INVITE
const inviteOK = "341"
