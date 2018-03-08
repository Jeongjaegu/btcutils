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
		fmt.Println("Encrypts the provided data (string) with the provided public key (hex encoded) " +
			"using the Bitcoin ECDSA curve.")
		fmt.Println("Usage: encrypter [datahex] [publichex]")
		return
	}
	data := os.Args[1]

	pub, err := hex.DecodeString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	ciphertext, err := Encrypt(pub, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(ciphertext))

	
}

func Encrypt(public []byte, message string) ([]byte, error) {
	pubKey, err := btcec.ParsePubKey(public, btcec.S256())
	if err != nil {
	    fmt.Println("error")
	    return nil, err
	}

	ciphertext, err := btcec.Encrypt(pubKey, []byte(message))
	if err != nil {
	    fmt.Println("error")
	    return nil, err
	}
	//fmt.Println(ciphertext)
	// n := bytes.IndexByte(ciphertext, 0)
	// cypherstring := string(ciphertext[:n])
	return ciphertext, nil
}

