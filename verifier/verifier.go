package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 4 {
		fmt.Println("verifies the provided signature with the provided public key (hex encoded) " +
			"using the Bitcoin ECDSA curve.")
		fmt.Println("Usage: verifier [signature] [data] [publichex]")
		return
	}

	sig, err := hex.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data, err := hex.DecodeString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	pubKeyBytes, err := hex.DecodeString(os.Args[3])
	if err != nil {
	    fmt.Println(err)
	    return
	}
	
	verified := Verify(sig, pubKeyBytes, data)
	fmt.Println(verified)
}

func Verify(sig, pubKeyBytes, messageHash []byte ) (bool) {		
	pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	if err != nil {
	    fmt.Println(err)
	    return false
	}

	signature, err := btcec.ParseSignature(sig, btcec.S256())
	if err != nil {
	    fmt.Println(err)
	    return false
	}
	// Verify the signature for the message using the public key.
	verified := signature.Verify(messageHash, pubKey)
	fmt.Println("Signature Verified?", verified)

	return verified
}

	
	