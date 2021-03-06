// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package sr25519

import (
	"crypto/rand"
	"strconv"

	"github.com/centrifuge/go-substrate-rpc-client/v3/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rjman-ljm/sherpax-utils/crypto"
)

var _ crypto.Keypair = &Keypair{}

type Keypair struct {
	keyringPair *signature.KeyringPair
}

func GenerateKeypair(network string) (*Keypair, error) {
	data := make([]byte, 32)
	_, err := rand.Read(data)
	if err != nil {
		return nil, err
	}
	return NewKeypairFromSeed("//"+hexutil.Encode(data), network)
}

func NewKeypairFromSeed(seed, network string) (*Keypair, error) {
	net, _ := strconv.ParseInt(network, 10, 8)
	kp, err := signature.KeyringPairFromSecret(seed, uint8(net))
	return &Keypair{&kp}, err
}

func NewKeypairFromKRP(pair signature.KeyringPair) *Keypair {
	return &Keypair{&pair}
}

// AsKeyringPair returns the underlying KeyringPair
func (kp *Keypair) AsKeyringPair() *signature.KeyringPair {
	return kp.keyringPair
}

// Encode uses scale to encode underlying KeyringPair
func (kp *Keypair) Encode() []byte {
	out, _ := types.EncodeToBytes(kp.keyringPair)
	return out
}

// Decode initializes keypair by decoding input as a KeyringPair
func (kp *Keypair) Decode(in []byte) error {
	kp.keyringPair = &signature.KeyringPair{}
	return types.DecodeFromBytes(in, kp.keyringPair)
}

// Address returns the ss58 formated address
func (kp *Keypair) Address() string {
	return kp.keyringPair.Address
}

// PublicKey returns the publickey encoded as a string
func (kp *Keypair) PublicKey() string {
	return hexutil.Encode(kp.keyringPair.PublicKey)
}
