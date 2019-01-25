package irc

const (
	Error   = "ERROR"
	Join    = "JOIN"
	Mode    = "MODE"
	Nick    = "NICK"
	Notice  = "NOTICE"
	Part    = "PART"
	Ping    = "PING"
	Privmsg = "PRIVMSG"
	Quit    = "QUIT"
	Topic   = "TOPIC"

	ReplyWelcome         = "001"
	ReplyYourHost        = "002"
	ReplyCreated         = "003"
	ReplyISupport        = "005"
	ReplyLUserClient     = "251"
	ReplyLUserOp         = "252"
	ReplyLUserUnknown    = "253"
	ReplyLUserChannels   = "254"
	ReplyLUserMe         = "255"
	ReplyAway            = "301"
	ReplyWhoisUser       = "311"
	ReplyWhoisServer     = "312"
	ReplyWhoisOperator   = "313"
	ReplyWhoisIdle       = "317"
	ReplyEndOfWhois      = "318"
	ReplyWhoisChannels   = "319"
	ReplyList            = "322"
	ReplyListEnd         = "323"
	ReplyNoTopic         = "331"
	ReplyTopic           = "332"
	ReplyNamReply        = "353"
	ReplyEndOfNames      = "366"
	ReplyMotd            = "372"
	ReplyMotdStart       = "375"
	ReplyEndOfMotd       = "376"
	ErrErroneousNickname = "432"
	ErrNicknameInUse     = "433"
	ErrForward           = "470"
)
