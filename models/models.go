package models

// Error ...
type Error struct {
	Error string `json:"error"`
}

// HostInfo comes from a node telling us it's alive
type HostInfo struct {
	Host   string `json:"host" binding:"required"`
	PubKey string `json:"pubKey" binding:"required"`
}

// MixHostInfo comes from a node telling us it's alive
type MixHostInfo struct {
	HostInfo
	Layer int `json:"layer" binding:"required"`
}

// MixNodePresence has presence info for a mixnode (including mixnode layer)
type MixNodePresence struct {
	MixHostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}

// Presence lets the server tell clients when a node was last seen
type Presence struct {
	HostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}
