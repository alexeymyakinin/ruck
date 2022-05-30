package model

type Message struct {
	ID        uint64
	Text      string
	AuthorID  uint64
	Timestamp uint64
}
