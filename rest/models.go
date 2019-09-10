package rest

// Error ...
type Error struct {
	Error string `json:"error"`
}

// ObjectRequest contains a full json graph
type ObjectRequest struct {
	Object interface{} `json:"object"`
}

// ObjectIDResponse contains a full json graph
type ObjectIDResponse struct {
	ID string `json:"id"`
}

}
