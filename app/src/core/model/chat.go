package model

type (
	Chat struct {
		ID   uint64
		Name string
	}

	ChatUser struct {
		ID      uint64
		Chat    uint64
		Account uint64
	}

	ChatMessage struct {
		ID        uint64
		Text      string
		Chat      uint64
		Sender    uint64
		Timestamp uint64
	}
)
