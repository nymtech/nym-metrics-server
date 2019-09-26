package models

// CocoHostInfo comes from a coconut node telling us it's alive
type CocoHostInfo struct {
	HostInfo
	Type string `json:"type" binding:"required"`
}

// CocoPresence holds presence info for a coconut node.
type CocoPresence struct {
	CocoHostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}

// HostInfo comes from a node telling us it's alive
type HostInfo struct {
	Host   string `json:"host"`
	PubKey string `json:"pubKey" binding:"required"`
}

// MixProviderHostInfo comes from a node telling us it's alive
type MixProviderHostInfo struct {
	HostInfo
	RegisteredClients []RegisteredClient `json:"registeredClients" binding:"required"`
}

// MixProviderPresence holds presence info for a mix provider node
type MixProviderPresence struct {
	MixProviderHostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}

// MixNodePresence holds presence info for a mixnode
type MixNodePresence struct {
	MixHostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}

// MixHostInfo comes from a node telling us it's alive
type MixHostInfo struct {
	HostInfo
	Layer uint `json:"layer" binding:"required"`
}

// Presence lets the server tell clients when a node was last seen
type Presence struct {
	HostInfo
	LastSeen int64 `json:"lastSeen" binding:"required"`
}

// RegisteredClient holds information about client registered at a provider
type RegisteredClient struct {
	PubKey string `json:"pubKey" binding:"required"`
}

// Topology shows us the current state of the overall Nym network
type Topology struct {
	CocoNodes        []CocoPresence
	MixNodes         []MixNodePresence
	MixProviderNodes []MixProviderPresence
}
