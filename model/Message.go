package model

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (m *Message) String() string {
	return m.Author + "пернул" + m.Body
}
