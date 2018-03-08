package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Encrypts the provided data (hex encoded) with the provided public key (hex encoded) " +
			"using the Bitcoin ECDSA curve.")
		fmt.Println("Usage: signer [datahex] [publichex]")
		return
	}
	data, err := hex.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	pub, err := hex.DecodeString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	ciphertext, err := Encrypt(pub, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ciphertext)
}

func Encrypt(public, message []byte) ([]byte, error) {
	pubKey, err := btcec.ParsePubKey(public, btcec.S256())
	if err != nil {
	    fmt.Println(err)
	    return nil, err
	}

	ciphertext, err := btcec.Encrypt(pubKey, []byte(message))
	if err != nil {
	    fmt.Println(err)
	    return nil, err
	}

	return ciphertext, nil
}