package tgbot

import "errors"

type Message struct {
	Priority string
	Text     string
}

func NewMessage(pr string, txt string) *Message {
	return &Message{pr, txt}
}

type MessageDB struct {
	MDB []Message
}

func (mdb *MessageDB) Insert(m Message) error {
	mdb.MDB = append(mdb.MDB, m)
	return nil
}

func (mdb *MessageDB) Delete(m Message) error {
	for i, val := range mdb.MDB {
		if val.Priority == m.Priority && val.Text == m.Text {
			mdb.MDB = append(mdb.MDB[:i], mdb.MDB[i+1:]...)
			return nil
		}
	}
	return errors.New("no such message found")
}
