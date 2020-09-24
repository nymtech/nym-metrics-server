package fixtures

import "github.com/nymtech/nym-directory/models"

// MixStatusesList A list of mix statuses
func MixStatusesList() []models.PersistedMixStatus {
	m1 := models.PersistedMixStatus{
		MixStatus: models.MixStatus{
			IPVersion: "6",
			PubKey:    "pubkey1",
			Up:        true,
		},
		Timestamp: 123,
	}

	m2 := models.PersistedMixStatus{
		MixStatus: models.MixStatus{
			IPVersion: "6",
			PubKey:    "pubkey1",
			Up:        true,
		},
		Timestamp: 1234,
	}

	statuses := []models.PersistedMixStatus{m1, m2}
	return statuses
}

// XSSMixStatus ...
func XSSMixStatus() models.MixStatus {
	xss := models.MixStatus{
		IPVersion: "6",
		PubKey:    "pubkey2<script>alert('gotcha')</script>",
		Up:        true,
	}
	return xss
}

// GoodMixStatus ...
func GoodMixStatus() models.MixStatus {
	return models.MixStatus{
		IPVersion: "6",
		PubKey:    "pubkey2",
		Up:        true,
	}
}

// GoodPersistedMixStatus ...
func GoodPersistedMixStatus() models.PersistedMixStatus {
	return models.PersistedMixStatus{
		MixStatus: GoodMixStatus(),
		Timestamp: 1234,
	}
}
