package models

import (
	_ "github.com/jinzhu/gorm"
)

// MixStatus indicates whether a given node is up or down, as reported by a Nym monitor node.
type MixStatus struct {
	PubKey    string `json:"pubKey" binding:"required" gorm:"index"`
	IPVersion string `json:"ipVersion" binding:"required"`
	Up        *bool  `json:"up" binding:"required"`
}

// PersistedMixStatus is a saved MixStatus with a timestamp recording when it
// was seen by the directory server. It can be used to build visualizations of
// mixnode uptime.
type PersistedMixStatus struct {
	MixStatus
	Timestamp int64 `json:"timestamp" binding:"required"`
}

// MixStatusReport gives a quick view of mixnode uptime performance
type MixStatusReport struct {
	PubKey   string
	LastIpv4 PersistedMixStatus
	LastIpv6 PersistedMixStatus
	// Last5Ipv4   []PersistedMixStatus
	// Last5Ipv6   []PersistedMixStatus
	// Last100Ipv4 []PersistedMixStatus
	// Last100Ipv6 []PersistedMixStatus
}
