/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:54
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:25:12
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

/**
 * Creates a new Asset instance by converting the given
 * context string, mobileID string and refeference string
 * into an abstract pathname. If the given strings are the
 * empty string, then the result is the empty Asset.
 *
 * @param		context 			string		Context string
 *					mobileID 			string		Mobile identification string
 *					reference 		string		A reference string
 * @return								Asset			A Asset struct
 */
func CreateAsset(context string, mobileID string, ref string) *Asset {
	// Generate a UUID
	id := uuid.New()
	// Create a new hash message with the data received
	asset := Asset{ID: id.String(), Owner: mobileID, Data: context, Reference: ref}
	return &asset
}

/**
 * Compress a list of Asset instances by converting the given
 * then into an string pathname. If the given array is empty
 * then the result is the empty string.
 *
 * @param		segments  []Asset		A list of Asset
 * @return						Asset			A Asset string
 */
func CompressAsset(segments []Asset) string {
	return segments[0].Data + "," + segments[len(segments)-1].Data
}

/**
 * Convert a Asset string into an hash. If the given array is empty
 * then the result is the empty hash.
 *
 * @param		data  	string		A Asset string
 * @return					string		A hash string
 */
func HashAsset(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
