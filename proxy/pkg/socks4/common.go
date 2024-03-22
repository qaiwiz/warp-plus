// Command is a SOCKS Command.
type Command byte

// String returns a human-readable representation of the SOCKS command.
func (cmd Command) String() string {
	switch cmd {
	case ConnectCommand:
		return "socks connect" // The command for establishing a connection to a remote server.
	default:
		return "socks " + strconv.Itoa(int(cmd))
	}
}

// reply is a SOCKS Command reply code.
type reply byte

// String returns a human-readable representation of the SOCKS reply code.
func (code reply) String() string {
	switch code {
	case grantedReply:
		return "request granted" // The request was successful.
	case rejectedReply:
		return "request rejected or failed" // The request was rejected or failed for an unspecified reason.
	case noIdentdReply:
		return "request rejected because SOCKS server cannot connect to identd on the client" // The SOCKS server cannot connect to the identd daemon on the client.
	case invalidUserReply:
		return "request rejected because the client program and identd report different user-ids" // The user-id reported by the client program and the identd daemon do not match.
	default:
		return "unknown code: " + strconv.Itoa(int(code)) // A reply code that is not defined in the SOCKS4 specification.
	}
}
