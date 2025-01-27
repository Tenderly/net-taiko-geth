// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/tenderly/net-taiko-geth/common"
)

// MarshalJSON marshals as JSON.
func (a AccessTuple) MarshalJSON() ([]byte, error) {
	type AccessTuple struct {
		Address     common.Address `json:"address"     gencodec:"required"`
		StorageKeys []common.Hash  `json:"storageKeys" gencodec:"required"`
	}
	var enc AccessTuple
	enc.Address = a.Address
	enc.StorageKeys = a.StorageKeys
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (a *AccessTuple) UnmarshalJSON(input []byte) error {
	type AccessTuple struct {
		Address     *common.Address `json:"address"     gencodec:"required"`
		StorageKeys []common.Hash   `json:"storageKeys" gencodec:"required"`
	}
	var dec AccessTuple
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address == nil {
		return errors.New("missing required field 'address' for AccessTuple")
	}
	a.Address = *dec.Address
	if dec.StorageKeys == nil {
		return errors.New("missing required field 'storageKeys' for AccessTuple")
	}
	a.StorageKeys = dec.StorageKeys
	return nil
}
