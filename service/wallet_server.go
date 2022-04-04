package service

import (
	"context"
	"encoding/csv"
	"encoding/hex"
	"strings"

	"github.com/berkeozerr/GoWalletApi/pb"
	btcutils "github.com/berkeozerr/GoWalletApi/thirdparty"
	"github.com/keyfuse/tokucore/network"
	"github.com/keyfuse/tokucore/xcore"
	bip32v2 "github.com/keyfuse/tokucore/xcore/bip32"
	"github.com/prettymuchbryce/hellobitcoin/base58check"
	"github.com/tyler-smith/go-bip39"
)

// Wallet server is server for Wallet Services

type BitcoinWalletServiceServer struct {
	pb.UnimplementedBitcoinWalletServiceServer
}

// Function generates new wallet server
func NewWalletServer() *BitcoinWalletServiceServer {
	return &BitcoinWalletServiceServer{}
}

//Generating random Mnemonic function
func (server *BitcoinWalletServiceServer) GenerateRandomMnemonic(context.Context, *pb.RandomMnemonicRequest) (*pb.RandomMnemonicResponse, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	response := &pb.RandomMnemonicResponse{
		Mnemonic: mnemonic,
	}
	return response, nil
}

//Generating HDSegWitAddress function
func (server *BitcoinWalletServiceServer) GenerateHDSegWitAddress(ctx context.Context, req *pb.GenerateHDSegWitAddressRequest) (*pb.GenerateHDSegWitAddressResponse, error) {
	seed, err := hex.DecodeString("0x" + req.GetSeed())
	if err != nil {
		return nil, err
	}
	hdKey := bip32v2.NewHDKey(seed)
	prvKey, err := hdKey.DeriveByPath(req.GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := prvKey.PublicKey()
	addr := xcore.NewPayToWitnessV0PubKeyHashAddress(publicKey.Hash160())
	response := &pb.GenerateHDSegWitAddressResponse{
		GeneratedAddress: addr.ToString(network.MainNet),
	}
	return response, nil

}

//Generating MultiSigP2SHAddress function
func (server *BitcoinWalletServiceServer) GenerateMultiSigP2SHAddress(ctx context.Context, req *pb.GenerateMultiSigP2SHAddressRequest) (*pb.GenerateMultiSigP2SHAddressResponse, error) {
	publicKeysString := strings.Join(req.PublicKeys[:], ",")
	address, _, err := GenerateAddress(int(req.GetM()), int(req.GetN()), publicKeysString)
	if err != nil {
		return nil, err
	}
	response := &pb.GenerateMultiSigP2SHAddressResponse{
		GeneratedAddress: address,
	}
	return response, nil
}

func GenerateAddress(flagM int, flagN int, flagPublicKeys string) (string, string, error) {
	//Convert public keys argument into slice of public key bytes with necessary tidying
	flagPublicKeys = strings.Replace(flagPublicKeys, "'", "\"", -1) //Replace single quotes with double since csv package only recognizes double quotes
	publicKeyStrings, err := csv.NewReader(strings.NewReader(flagPublicKeys)).Read()
	if err != nil {
		return "", "", err
	}
	publicKeys := make([][]byte, len(publicKeyStrings))
	for i, publicKeyString := range publicKeyStrings {
		publicKeyString = strings.TrimSpace(publicKeyString)   //Trim whitespace
		publicKeys[i], err = hex.DecodeString(publicKeyString) //Get private keys as slice of raw bytes
		if err != nil {
			return "", "", err
		}
	}
	//Create redeemScript from public keys
	redeemScript, err := btcutils.NewMOfNRedeemScript(flagM, flagN, publicKeys)
	if err != nil {
		return "", "", err
	}
	redeemScriptHash, err := btcutils.Hash160(redeemScript)
	if err != nil {
		return "", "", err
	}
	//Get P2SH address by base58 encoding with P2SH prefix 0x05
	P2SHAddress := base58check.Encode("05", redeemScriptHash)
	//Get redeemScript in Hex
	redeemScriptHex := hex.EncodeToString(redeemScript)

	return P2SHAddress, redeemScriptHex, nil
}
