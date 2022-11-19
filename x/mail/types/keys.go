package types

const (
	// ModuleName defines the module name
	ModuleName = "mail"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mail"

	// Version defines the current version the IBC module supports
	Version = "mail-1"

	// PortID is the default port id that module binds to
	PortID = "mail"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("mail-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MessageKey      = "Message/value/"
	MessageCountKey = "Message/count/"
)

const (
	SentMessageKey      = "SentMessage/value/"
	SentMessageCountKey = "SentMessage/count/"
)
