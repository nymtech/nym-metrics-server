package models

import "time"

// Error ...
type Error struct {
	Error string `json:"error"`
}

// ObjectRequest contains a full json graph
type ObjectRequest struct {
	Object interface{} `json:"object"`
}

// ObjectIDResponse contains a full json graph
type ObjectIDResponse struct {
	ID string `json:"id"`
}

// UpMsg comes from a node telling us it's alive
type UpMsg struct {
	PubKey string `json:"pubKey" binding:"required"`
}

// Presence lets the server tell clients wahen a node was last seen
type Presence struct {
	LastSeen time.Time `json:"lastSeen" binding:"required" time_format:"unix"`
	PubKey   string    `json:"pubKey" binding:"required"`
}
