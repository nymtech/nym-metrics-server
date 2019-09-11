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

// HostInfo comes from a node telling us it's alive
type HostInfo struct {
	Host   string `json:"host" binding:"required"`
	Layer  int    `json:"layer" binding:"required"`
	PubKey string `json:"pubKey" binding:"required"`
}

// Presence lets the server tell clients when a node was last seen
type Presence struct {
	HostInfo
	LastSeen time.Time `json:"lastSeen" binding:"required" time_format:"unix"`
}
