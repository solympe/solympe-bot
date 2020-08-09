package models

// Update ...
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Chat ...
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// UpdateResponse ...
type UpdateResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// BotMessage ...
type BotMessage struct {
	ID   int    `json:"chat_id"`
	Text string `json:"text"`
}
