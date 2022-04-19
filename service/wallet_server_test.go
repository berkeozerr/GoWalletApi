package service_test

import (
	"context"
	"testing"

	"github.com/berkeozerr/GoWalletApi/pb"
	"github.com/berkeozerr/GoWalletApi/service"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomMnemonic(t *testing.T) {
	req := &pb.RandomMnemonicRequest{}
	server := service.NewWalletServer()
	res, err := server.GenerateRandomMnemonic(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotEmpty(t, res.Mnemonic)

}
func TestGenerateHDSegWitAddress(t *testing.T) {

	req := &pb.GenerateHDSegWitAddressRequest{
		Seed: "lizard century nut catch figure build swift call pledge toe cereal truck recipe faint clerk",
		Path: "m/84'/0'/0'/0/0",
	}
	server := service.NewWalletServer()
	res, _ := server.GenerateHDSegWitAddress(context.Background(), req)
	require.Equal(t, res.GeneratedAddress, "bc1q0j5dewvk89ss00l68a9hf8lah4c66wmahddmdv")

}

func TestGenerateMultiSigP2SHAddress(t *testing.T) {
	req := &pb.GenerateMultiSigP2SHAddressRequest{}
	req.M = 2
	req.N = 3
	expectedResult := "36NUkt6FWUi3LAWBqWRdDmdTWbt91Yvfu7"
	req.PublicKeys = []string{"026477115981fe981a6918a6297d9803c4dc04f328f22041bedff886bbc2962e01", "02c96db2302d19b43d4c69368babace7854cc84eb9e061cde51cfa77ca4a22b8b9,03c6103b3b83e4a24a0e33a4df246ef11772f9992663db0c35759a5e2ebf68d8e9"}
	server := service.NewWalletServer()
	//	res, err := service.GenerateAddress(2, 3, testPublicKeys)
	res, _ := server.GenerateMultiSigP2SHAddress(context.Background(), req)
	//	fmt.Println(res, err)
	require.Equal(t, res.GeneratedAddress, expectedResult)
}
