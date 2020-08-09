package models

// UpdateResponse ...
type UpdateResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update ...
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}
