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
		fmt.Println("decrypts the provided data (hex encoded) with the provided private key (hex encoded) " +
			"using the Bitcoin ECDSA curve.")
		fmt.Println("Usage: decrypter [cipher] [privatehex]")
		return
	}
	 
	cipher, err := hex.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	priv, err := hex.DecodeString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	message, err := Decrypt(priv,
						(cipher))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(message))
}


func Decrypt(pkBytes, cipher []byte) ([]byte, error) {
	// Decode the hex-encoded private key.
	

	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	// ciphertext, err := hex.DecodeString(cypherstring)

	// Try decrypting the message.
	plaintext, err := btcec.Decrypt(privKey, (cipher))
	if err != nil {
	    fmt.Println(err)
	    return nil, nil
	}

	return plaintext, nil

}