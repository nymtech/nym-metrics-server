package fixtures

import "github.com/nymtech/nym-directory/models"

// MixStatusesList A list of mix statuses
func MixStatusesList() []models.PersistedMixStatus {
	upTrue := true
	m1 := models.PersistedMixStatus{
		MixStatus: models.MixStatus{
			IPVersion: "6",
			PubKey:    "pubkey1",
			Up:        &upTrue,
		},
		Timestamp: 123,
	}

	m2 := models.PersistedMixStatus{
		MixStatus: models.MixStatus{
			IPVersion: "6",
			PubKey:    "pubkey1",
			Up:        &upTrue,
		},
		Timestamp: 1234,
	}

	statuses := []models.PersistedMixStatus{m1, m2}
	return statuses
}

// XSSMixStatus ...
func XSSMixStatus() models.MixStatus {
	upTrue := true
	xss := models.MixStatus{
		IPVersion: "6",
		PubKey:    "pubkey2<script>alert('gotcha')</script>",
		Up:        &upTrue,
	}
	return xss
}

// GoodMixStatus ...
func GoodMixStatus() models.MixStatus {
	upTrue := true
	return models.MixStatus{
		IPVersion: "6",
		PubKey:    "pubkey2",
		Up:        &upTrue,
	}
}

// GoodPersistedMixStatus ...
func GoodPersistedMixStatus() models.PersistedMixStatus {
	return models.PersistedMixStatus{
		MixStatus: GoodMixStatus(),
		Timestamp: 1234,
	}
}
