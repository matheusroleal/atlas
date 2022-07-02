/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:54
 * @Last Modified by:   Matheus Leal
 * @Last Modified time: 2022-07-01 22:54:54
 */
package asset

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/google/uuid"
)

type Asset struct {
	ID        string `json:"ID"`
	Owner     string `json:"owner"`
	Data      string `json:"data"`
	Reference string `json:"reference"`
}

func CreateAsset(context string, mobileID string, ref string) *Asset {
	// Generate a UUID
	id := uuid.New()
	// Create a new hash message with the data received
	asset := Asset{ID: id.String(), Owner: mobileID, Data: context, Reference: ref}
	return &asset
}

func CompressAsset(segments []Asset) string {
	return segments[0].Data + "," + segments[len(segments)-1].Data
}

func HashAsset(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
