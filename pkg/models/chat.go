package models

// Chat ...
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// BotMessage ...
type BotMessage struct {
	ID   int    `json:"chat_id"`
	Text string `json:"text"`
}
