package models

// MixStatus indicates whether a given node is up or down, as reported by a Nym monitor node.
type MixStatus struct {
	PubKey    string `json:"pubKey" binding:"required"`
	IPVersion string `json:"protocol" binding:"required"`
	Up        bool   `json:"status" binding:"required"`
}

// PersistedMixStatus is a saved MixStatus with a timestamp recording when it
// was seen by the directory server. It can be used to build visualizations of
// mixnode uptime.
type PersistedMixStatus struct {
	MixStatus
	Timestamp int64 `json:"timestamp" binding:"required"`
}
