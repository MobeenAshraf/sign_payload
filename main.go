package main

import (
	"encoding/hex"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/keys"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func main() {
	privKeyHex := ""
	senderAddress := ""
	signingMessage := ""
	const (
		payloadSigType types.SignatureType = "ecdsa_recovery"
		sigType        types.SignatureType = "ecdsa_recovery"
	)

	signingPayloadHexDecoded, _ := hex.DecodeString(signingMessage)
	keyPair, err := keys.ImportPrivateKey(privKeyHex, types.Secp256k1)
	if err != nil {
		returnErr(err)
		return
	}
	signer, _ := keyPair.Signer()
	sig, err := signer.Sign(&types.SigningPayload{
		AccountIdentifier: &types.AccountIdentifier{
			Address: senderAddress,
		},
		Bytes:         signingPayloadHexDecoded,
		SignatureType: payloadSigType,
	}, sigType)
	if err != nil {
		returnErr(err)
		return
	}
	fmt.Println(hex.EncodeToString(sig.Bytes))

}

func returnErr(err error) {
	fmt.Println("Error:")
	fmt.Println(err)
}
