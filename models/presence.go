package models

// HostInfo comes from a node telling us it's alive
type HostInfo struct {
	Host   string `json:"host" binding:"required"`
	PubKey string `json:"pubKey" binding:"required"`
}

// MixNodePresence holds presence info for a mixnode
type MixNodePresence struct {
	MixHostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}

// Presence lets the server tell clients when a node was last seen
type Presence struct {
	HostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}

// MixHostInfo comes from a node telling us it's alive
type MixHostInfo struct {
	HostInfo
	Layer uint `json:"layer" binding:"required"`
}
