package asset

import (
  "github.com/google/uuid"
)

type Asset struct {
	ID			    string	`json:"ID"`
	Owner		    string	`json:"owner"`
	Checkpoint	string	`json:"checkpoint"`
}

func CreateCheckpoint(data string, mobileID string) *Asset {
  // Generate a UUID
  id := uuid.New()
  // Create a new Checkpoint with the data received
	asset := Asset{ ID: id.String(), Owner: mobileID, Checkpoint: data }
  return &asset
}
