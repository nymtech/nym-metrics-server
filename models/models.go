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
	Layer uint `json:"layer" binding:"required"`
}

// MixMetric is a report from each mixnode detailing recent traffic.
// Useful for creating visualisations.
type MixMetric struct {
	PubKey   string          `json:"pubKey" binding:"required"`
	Sent     map[string]uint `json:"sent" binding:"required"`
	Received uint            `json:"received" binding:"required"`
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
