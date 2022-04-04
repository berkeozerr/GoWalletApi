package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/berkeozerr/GoWalletApi/pb"
	"github.com/berkeozerr/GoWalletApi/service"
)

func TestGenerateRandomMnemonic(t *testing.T) {
	req := &pb.RandomMnemonicRequest{}
	server := service.NewWalletServer()
	res, err := server.GenerateRandomMnemonic(context.Background(), req)
	fmt.Println(res, err)
	if err != nil {
		t.Fatal(err)
	}

}
func TestGenerateHDSegWitAddress(t *testing.T) {
	req := &pb.GenerateHDSegWitAddressRequest{}
	req.SeedPhrase = "lizard century nut catch figure build swift call pledge toe cereal truck recipe faint clerk"
	req.Path = "m/84'/0'/0'/0/0"
	server := service.NewWalletServer()
	res, err := server.GenerateHDSegWitAddress(context.Background(), req)
	fmt.Println(res, err)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGenerateMultiSigP2SHAddress(t *testing.T) {
	req := &pb.GenerateMultiSigP2SHAddressRequest{}
	req.M = 2
	req.N = 3
	testPublicKeys := "04a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd,046ce31db9bdd543e72fe3039a1f1c047dab87037c36a669ff90e28da1848f640de68c2fe913d363a51154a0c62d7adea1b822d05035077418267b1a1379790187,0411ffd36c70776538d079fbae117dc38effafb33304af83ce4894589747aee1ef992f63280567f52f5ba870678b4ab4ff6c8ea600bd217870a8b4f1f09f3a8e83"
	//testPublicKeys := "026477115981fe981a6918a6297d9803c4dc04f328f22041bedff886bbc2962e01,02c96db2302d19b43d4c69368babace7854cc84eb9e061cde51cfa77ca4a22b8b9,03c6103b3b83e4a24a0e33a4df246ef11772f9992663db0c35759a5e2ebf68d8e9"
	//server := service.NewWalletServer()
	res, err := service.GenerateAddress(2, 3, testPublicKeys)
	//	res, err := server.GenerateMultiSigP2SHAddress(context.Background(), req)
	fmt.Println(res, err)
	/*	if err != nil {
		t.Fatal(err)
	}*/

}
