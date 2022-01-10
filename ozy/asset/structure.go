package asset

import (
	"github.com/google/uuid"
)

type Asset struct {
	ID    string `json:"ID"`
	Owner string `json:"owner"`
	Data  string `json:"data"`
}

func CreateHashData(data string, mobileID string) *Asset {
	// Generate a UUID
	id := uuid.New()
	// Create a new hash message with the data received
	asset := Asset{ID: id.String(), Owner: mobileID, Data: data}
	return &asset
}
