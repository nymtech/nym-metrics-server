package fixtures

import "github.com/nymtech/nym-directory/models"

func Disallow() models.Disallow {
	d := models.Disallow{
		HostKey: "abc123",
	}
	return d
}

func GoodCocoHost() models.CocoHostInfo {
	good := models.CocoHostInfo{
		HostInfo: models.HostInfo{
			Location: "foomplandia",
			Host:     ":1789",
			PubKey:   "pubkey",
			Version:  "0.1.0",
		},
		Type: "type",
	}
	return good
}

func GoodMixHost() models.MixHostInfo {
	good := models.MixHostInfo{
		HostInfo: models.HostInfo{
			Location: "foomplandia",
			Host:     ":1789",
			PubKey:   "pubkey",
			Version:  "0.1.0",
		},
		Layer: 1,
	}
	return good
}

func GoodMixProviderHost() models.MixProviderHostInfo {
	client1 := models.RegisteredClient{PubKey: "client1"}
	client2 := models.RegisteredClient{PubKey: "client2"}
	clients := []models.RegisteredClient{client1, client2}
	good := models.MixProviderHostInfo{
		Location:          "foomplandia",
		ClientListener:    ":1789",
		MixnetListener:    ":1790",
		PubKey:            "pubkey",
		Version:           "0.1.0",
		RegisteredClients: clients,
	}
	return good
}

func XssCocoHost() models.CocoHostInfo {
	xss := models.CocoHostInfo{
		HostInfo: models.HostInfo{
			Location: "foomplandia",
			Host:     ":1789",
			PubKey:   "pubkey<script>alert('gotcha')</script>",
			Version:  "0.1.0",
		},
		Type: "type<script>alert('gotcha')",
	}
	return xss
}

func XssMixHost() models.MixHostInfo {
	xss := models.MixHostInfo{
		HostInfo: models.HostInfo{
			Location: "foomplandia",
			Host:     ":1789",
			PubKey:   "pubkey<script>alert('gotcha')</script>",
			Version:  "0.1.0",
		},
		Layer: 1,
	}
	return xss
}

func XssMixProviderHost() models.MixProviderHostInfo {
	client1 := models.RegisteredClient{PubKey: "client1<script>alert('gotcha')</script>"}
	client2 := models.RegisteredClient{PubKey: "client2<script>alert('gotcha')</script>"}
	clients := []models.RegisteredClient{client1, client2}
	xss := models.MixProviderHostInfo{
		Location:          "foomplandia",
		ClientListener:    ":1789",
		MixnetListener:    ":1790",
		PubKey:            "pubkey<script>alert('gotcha')</script>",
		Version:           "0.1.0",
		RegisteredClients: clients,
	}
	return xss
}
