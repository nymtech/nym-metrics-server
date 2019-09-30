package presence

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/nymtech/nym-directory/models"
)

// Sanitizer cleans untrusted input fields
type Sanitizer interface {
	Sanitize(interface{})
}

// NewCoconodeSanitizer constructor...
func NewCoconodeSanitizer(p *bluemonday.Policy) CoconodeSanitizer {
	return CoconodeSanitizer{
		policy: p,
	}
}

// CoconodeSanitizer kills untrusted input in CocoHostInfo structs
type CoconodeSanitizer struct {
	policy *bluemonday.Policy
}

// Sanitize CocoHostInfo input
func (s *CoconodeSanitizer) Sanitize(input models.CocoHostInfo) models.CocoHostInfo {
	input.PubKey = s.policy.Sanitize(input.PubKey)
	input.Host = s.policy.Sanitize(input.Host)
	input.Type = s.policy.Sanitize(input.Type)
	return input
}

// NewMixnodeSanitizer constructor...
func NewMixnodeSanitizer(p *bluemonday.Policy) MixnodeSanitizer {
	return MixnodeSanitizer{
		policy: p,
	}
}

// MixnodeSanitizer kills untrusted input in MixHostInfo structs
type MixnodeSanitizer struct {
	policy *bluemonday.Policy
}

// Sanitize MixHostInfo input
func (s *MixnodeSanitizer) Sanitize(input models.MixHostInfo) models.MixHostInfo {
	input.PubKey = s.policy.Sanitize(input.PubKey)
	input.Host = s.policy.Sanitize(input.Host)
	return input
}

// MixproviderSanitizer kills untrusted input in MixProviderHostInfo structs
type MixproviderSanitizer struct {
	policy *bluemonday.Policy
}
