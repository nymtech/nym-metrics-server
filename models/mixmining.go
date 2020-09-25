package models

import (
	_ "github.com/jinzhu/gorm"
)

// MixStatus indicates whether a given node is up or down, as reported by a Nym monitor node.
type MixStatus struct {
	PubKey    string `json:"pubKey" binding:"required" gorm:"index"`
	IPVersion string `json:"ipVersion" binding:"required"`
	Up        bool   `json:"up" binding:"required"`
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
	PubKey     string
	IPV4Status UptimeReport
	IPV6Status UptimeReport
}

// UptimeReport models an uptime report for IPV4 or IPV6 protocol
type UptimeReport struct {
	IPVersion    string
	MostRecent   bool
	Last5Minutes int
	LastHour     int
	LastDay      int
	LastWeek     int
	LastMonth    int
}
