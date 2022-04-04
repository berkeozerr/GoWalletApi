// Package btcutils contains a number of useful Bitcoin related functions originally used in the go-bitcoin-multisig
// project but useful in any general Bitcoin project.
package btcutils

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	mathrand "math/rand"
	"time"

	"golang.org/x/crypto/ripemd160"
)

// setFixedNonce is used for testing and debugging. It is by default false, but if set to true, then newNonce()
// will always return a zero-valued [32]byte{}. Allows repeatable ECDSA signatures for testing
// **Should never be turned on in production. Limit to use in tests only.**
var SetFixedNonce bool

//Fixed nonce value for repeatable testing.
//We declare var and not const because Go slices are mutable and cannot be const, but we use fixedNonce like a constant.
var FIXED_NONCE = [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

func randInt(min int, max int) uint8 {
	//This function is *not* cryptographically secure.
	//Used for nonce which does not need to be cryptographically secure, just changing with each signature
	mathrand.Seed(time.Now().UTC().UnixNano())
	return uint8(min + mathrand.Intn(max-min))
}

func newNonce() *[32]byte {
	//This function is *not* cryptographically secure.
	//Used for nonce which does not need to be cryptographically secure, just changing with each signature
	var bytes [32]byte
	if !SetFixedNonce {
		for i := 0; i < 32; i++ {
			bytes[i] = byte(randInt(0, math.MaxUint8))
		}
	} else {
		bytes = FIXED_NONCE
	}
	return &bytes
}

// NewRandomBytes generates pseudorandom bytes of length size.
// Cryptographically secure to the limits of crypto/rand package.
func NewRandomBytes(size int) ([]byte, error) {
	randBytes := make([]byte, size)
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, err
	}
	return randBytes, nil
}

// Hash160 performs the same operations as OP_HASH160 in Bitcoin Script
// It hashes the given data first with SHA256, then RIPEMD160
func Hash160(data []byte) ([]byte, error) {
	//Does identical function to Script OP_HASH160. Hash once with SHA-256, then RIPEMD-160
	if data == nil {
		return nil, errors.New("Empty bytes cannot be hashed")
	}
	shaHash := sha256.New()
	shaHash.Write(data)
	hash := shaHash.Sum(nil) //SHA256 first
	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(hash)
	hash = ripemd160Hash.Sum(nil) //RIPEMD160 second

	return hash, nil
}

// NewMOfNRedeemScript creates a M-of-N Multisig redeem script given m, n and n public keys
func NewMOfNRedeemScript(m int, n int, publicKeys [][]byte) ([]byte, error) {
	//Check we have valid numbers for M and N
	if n < 1 || n > 7 {
		return nil, errors.New("N must be between 1 and 7 (inclusive) for valid, standard P2SH multisig transaction as per Bitcoin protocol.")
	}
	if m < 1 || m > n {
		return nil, errors.New("M must be between 1 and N (inclusive).")
	}
	//Check we have N public keys as necessary.
	if len(publicKeys) != n {
		return nil, errors.New(fmt.Sprintf("Need exactly %d public keys to create P2SH address for %d-of-%d multisig transaction. Only %d keys provided.", n, m, n, len(publicKeys)))
	}
	//Get OP Code for m and n.
	//81 is OP_1, 82 is OP_2 etc.
	//80 is not a valid OP_Code, so we floor at 81
	mOPCode := OP_1 + (m - 1)
	nOPCode := OP_1 + (n - 1)
	//Multisig redeemScript format:
	//<OP_m> <A pubkey> <B pubkey> <C pubkey>... <OP_n> OP_CHECKMULTISIG
	var redeemScript bytes.Buffer
	redeemScript.WriteByte(byte(mOPCode)) //m
	for _, publicKey := range publicKeys {
		err := CheckPublicKeyIsValid(publicKey)
		if err != nil {
			return nil, err
		}
		redeemScript.WriteByte(byte(len(publicKey))) //PUSH
		redeemScript.Write(publicKey)                //<pubkey>
	}
	redeemScript.WriteByte(byte(nOPCode)) //n
	redeemScript.WriteByte(byte(OP_CHECKMULTISIG))
	return redeemScript.Bytes(), nil
}

// CheckPublicKeyIsValid runs a couple of checks to make sure a public key looks valid.
// Returns an error with a helpful message or nil if key is valid.
func CheckPublicKeyIsValid(publicKey []byte) error {
	errMessage := ""
	if publicKey == nil {
		errMessage += "Public key cannot be empty.\n"
	} else if len(publicKey) != 33 {
		errMessage += fmt.Sprintf("Public key should be 33 bytes long. Provided public key is %d bytes long.", len(publicKey))
	} else if publicKey[0] != byte(2) && publicKey[0] != byte(3) {
		errMessage += fmt.Sprintf("Public key is uncompressed. Please provide compressed public key", hex.EncodeToString([]byte{publicKey[0]}))
	}
	if errMessage != "" {
		errMessage += "Invalid public key:\n"
		errMessage += hex.EncodeToString(publicKey)
		return errors.New(errMessage)
	}
	return nil
}
