package service

import (
	"context"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"log"
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

type RandomMnemonicGeneratorServer struct {
	pb.UnimplementedRandomMnemonicGeneratorServer
}

// Function generates new wallet server
func NewWalletServer() *RandomMnemonicGeneratorServer {
	return &RandomMnemonicGeneratorServer{}
}

//Generating random Mnemonic function
func (server *RandomMnemonicGeneratorServer) GenerateRandomMnemonic(context.Context, *pb.RandomMnemonicRequest) (*pb.RandomMnemonicResponse, error) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println("Mnemonic: ", mnemonic)
	response := &pb.RandomMnemonicResponse{
		Mnemonic: mnemonic,
	}
	return response, nil
}

//Generating HDSegWitAddress
func (server *RandomMnemonicGeneratorServer) GenerateHDSegWitAddress(ctx context.Context, req *pb.GenerateHDSegWitAddressRequest) (*pb.GenerateHDSegWitAddressResponse, error) {
	seed := bip39.NewSeed(req.GetSeedPhrase(), "")
	hdkey := bip32v2.NewHDKey([]byte(seed))
	prvkey, err := hdkey.DeriveByPath(req.GetPath())
	if err != nil {
		panic(err)
	}
	public_key := prvkey.PublicKey()
	addr := xcore.NewPayToWitnessV0PubKeyHashAddress(public_key.Hash160())
	response := &pb.GenerateHDSegWitAddressResponse{
		GeneratedAddress: addr.ToString(network.MainNet),
	}
	return response, nil

}
func (server *RandomMnemonicGeneratorServer) GenerateMultiSigP2SHAddress(ctx context.Context, req *pb.GenerateMultiSigP2SHAddressRequest) (*pb.GenerateMultiSigP2SHAddressResponse, error) {
	public_keys_string := strings.Join(req.PublicKeys[:], ",")
	address, _ := GenerateAddress(int(req.GetM()), int(req.GetN()), public_keys_string)
	response := &pb.GenerateMultiSigP2SHAddressResponse{
		GeneratedAddress: address,
	}
	return response, nil
}

func GenerateAddress(flagM int, flagN int, flagPublicKeys string) (string, string) {
	//Convert public keys argument into slice of public key bytes with necessary tidying
	flagPublicKeys = strings.Replace(flagPublicKeys, "'", "\"", -1) //Replace single quotes with double since csv package only recognizes double quotes
	publicKeyStrings, err := csv.NewReader(strings.NewReader(flagPublicKeys)).Read()
	if err != nil {
		log.Fatal(err)
	}
	publicKeys := make([][]byte, len(publicKeyStrings))
	for i, publicKeyString := range publicKeyStrings {
		publicKeyString = strings.TrimSpace(publicKeyString)   //Trim whitespace
		publicKeys[i], err = hex.DecodeString(publicKeyString) //Get private keys as slice of raw bytes
		if err != nil {
			log.Fatal(err, "\n", "Offending publicKey: \n", publicKeyString)
		}
	}
	//Create redeemScript from public keys
	redeemScript, err := btcutils.NewMOfNRedeemScript(flagM, flagN, publicKeys)
	if err != nil {
		log.Fatal(err)
	}
	redeemScriptHash, err := btcutils.Hash160(redeemScript)
	if err != nil {
		log.Fatal(err)
	}
	//Get P2SH address by base58 encoding with P2SH prefix 0x05
	P2SHAddress := base58check.Encode("05", redeemScriptHash)
	//Get redeemScript in Hex
	redeemScriptHex := hex.EncodeToString(redeemScript)

	return P2SHAddress, redeemScriptHex
}
