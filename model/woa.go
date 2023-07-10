package model

type WoaMarkdown struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
}

type Markdown struct {
	Text string `json:"text"`
}
